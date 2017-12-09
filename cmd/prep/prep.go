package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/imports"
)

const importPath = "github.com/hexdigest/prep"

type (
	queryFinder struct {
		packageInfo    *types.Info
		prepInterfaces map[string]*types.Interface
		queries        []string
	}
)

func main() {
	var (
		sourcePackage = flag.String("f", "", "source package import path, i.e. github.com/my/package")
		withTests     = flag.Bool("t", false, "parse test files")
	)
	flag.Parse()

	if *sourcePackage == "" {
		flag.PrintDefaults()
		return
	}

	conf := loader.Config{
		TypeChecker: types.Config{
			FakeImportC:              true,
			DisableUnusedImportCheck: true,
		},
		TypeCheckFuncBodies: func(path string) bool {
			return strings.HasPrefix(path, *sourcePackage)
		},
	}

	if *withTests {
		conf.ImportWithTests(*sourcePackage)
	} else {
		conf.Import(*sourcePackage)
	}

	conf.Import(importPath)

	prog, err := conf.Load()
	if err != nil {
		log.Fatalf("prep: failed to load package: %v\n", err)
	}

	pkg := prog.Package(*sourcePackage)
	finder := &queryFinder{
		packageInfo:    &pkg.Info,
		prepInterfaces: findInterfaces(prog.Package(importPath)),
	}

	for _, file := range pkg.Files {
		ast.Walk(finder, file)
	}

	path, err := getPathToPackage(*sourcePackage)
	if err != nil {
		log.Fatalf("prep: %v", err)
	}

	outputFileName := filepath.Join(path, "prepared_statements.go")

	queries := uniqueStrings(finder.queries)

	if len(queries) == 0 {
		log.Fatalf("prep: no SQL queries found")
	}

	code, err := generateCode(pkg.Pkg.Name(), queries)
	if err != nil {
		log.Fatalf("prep: failed to generate code: %v", err)
	}

	file, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("prep: failed to create file: %v", err)
	}
	defer file.Close()

	if _, err := file.Write(code); err != nil {
		log.Fatalf("prep: failed to write generated code to the file: %v", err)
	}
}

func getPathToPackage(importPath string) (string, error) {
	p, err := build.Default.Import(importPath, "", build.FindOnly)
	if err != nil {
		return "", fmt.Errorf("failed to detect absolute path of the package %q: %v", importPath, err)
	}

	return filepath.Join(p.SrcRoot, p.ImportPath), nil
}

func generateCode(packageName string, queries []string) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	fmt.Fprintf(buf,
		"package %s\n\nvar prepStatements = []string{\n%s,\n}",
		packageName, strings.Join(queries, ",\n"))

	formatted, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		return formatted, fmt.Errorf("failed to format code %s: %v", buf.Bytes(), err)
	}

	return formatted, nil
}

// uniqueStrings returns a sorted slice of the unique strings
// from the given strings slice
func uniqueStrings(strings []string) []string {
	m := make(map[string]struct{})
	for _, s := range strings {
		m[s] = struct{}{}
	}

	unique := []string{}

	for s := range m {
		unique = append(unique, s)
	}

	sort.Strings(unique)
	return unique
}

// maps method name to the interface it implements
var methodImplements = map[string]string{
	"Prepare":  "preparer",
	"Exec":     "executer",
	"Query":    "querier",
	"QueryRow": "rowQuerier",

	"PrepareContext":  "preparerWithContext",
	"ExecContext":     "executerWithContext",
	"QueryContext":    "querierWithContext",
	"QueryRowContext": "rowQuerierWithContext",
}

// Visit implements ast.Visitor interface
func (f *queryFinder) Visit(node ast.Node) ast.Visitor {
	fCall, ok := node.(*ast.CallExpr)
	if !ok {
		return f
	}

	selector, ok := fCall.Fun.(*ast.SelectorExpr)
	if !ok {
		return f
	}

	interfaceName := methodImplements[selector.Sel.Name]
	if interfaceName == "" {
		return f
	}

	if !types.Implements(f.packageInfo.TypeOf(selector.X), f.prepInterfaces[interfaceName]) {
		return nil
	}

	var query string
	switch selector.Sel.Name {
	case "Prepare", "Exec", "Query", "QueryRow":
		query = f.processQuery(fCall.Args[0])
	case "PrepareContext", "ExecContext", "QueryContext", "QueryRowContext":
		query = f.processQuery(fCall.Args[1])
	}

	if query != "" {
		f.queries = append(f.queries, query)
	}

	return nil
}

// processQuery returns a string value of the expression if the
// expression is either a string literal or a string constant otherwise
// an empty string is returned
func (f *queryFinder) processQuery(queryArg ast.Expr) string {
	switch q := queryArg.(type) {
	case *ast.BasicLit:
		return q.Value
	case *ast.Ident:
		obj := f.packageInfo.ObjectOf(q)

		if constant, ok := obj.(*types.Const); ok {
			return constant.Val().ExactString()
		}
	}

	return ""
}
