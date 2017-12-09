package main

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/loader"
)

type interfaceFinder struct {
	packageInfo     *types.Info
	foundInterfaces map[string]*types.Interface
}

// Visit implements ast.Visitor interface
func (f *interfaceFinder) Visit(node ast.Node) ast.Visitor {
	typeSpec, ok := node.(*ast.TypeSpec)
	if !ok {
		return f
	}

	typeName := typeSpec.Name.Name
	typeType := f.packageInfo.TypeOf(typeSpec.Type)
	if i, ok := typeType.(*types.Interface); ok {
		f.foundInterfaces[typeName] = i
	}

	return nil
}

func findInterfaces(pkg *loader.PackageInfo) map[string]*types.Interface {
	finder := &interfaceFinder{
		packageInfo:     &pkg.Info,
		foundInterfaces: map[string]*types.Interface{},
	}

	for _, file := range pkg.Files {
		ast.Walk(finder, file)
	}

	return finder.foundInterfaces
}
