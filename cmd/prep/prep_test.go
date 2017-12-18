package main

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"reflect"
	"testing"
)

func Test_queryFinder_processQuery(t *testing.T) {
	ident := &ast.Ident{}
	tests := []struct {
		name     string
		qf       *queryFinder
		queryArg ast.Expr
		want     string
	}{
		{
			name:     "query is a string literal",
			queryArg: &ast.BasicLit{Value: "string literal"},
			want:     "string literal",
		},
		{
			name:     "query is not a constant",
			queryArg: &ast.Ident{},
			qf: &queryFinder{
				packageInfo: &types.Info{
					Defs: map[*ast.Ident]types.Object{},
					Uses: map[*ast.Ident]types.Object{},
				},
			},
		},
		{
			name:     "query is a constant",
			queryArg: ident,
			qf: &queryFinder{
				packageInfo: &types.Info{
					Defs: map[*ast.Ident]types.Object{
						ident: types.NewConst(token.Pos(0), nil, "unnamed", &types.Basic{}, constant.MakeString("query")),
					},
				},
			},
			want: `"query"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.qf.processQuery(tt.queryArg); got != tt.want {
				t.Errorf("queryFinder.processQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryFinder_Visit(t *testing.T) {
	tests := []struct {
		name string
		qf   *queryFinder
		node ast.Node
		want ast.Visitor
	}{
		{
			name: "node is not a call expression",
			qf:   &queryFinder{},
			want: &queryFinder{},
		},
		{
			name: "function is called without selector",
			qf:   &queryFinder{},
			node: &ast.CallExpr{},
			want: &queryFinder{},
		},
		{
			name: "method name doesn't match",
			qf:   &queryFinder{packageInfo: &types.Info{}},
			node: &ast.CallExpr{Fun: &ast.SelectorExpr{Sel: &ast.Ident{Name: "WrongMethod"}}},
			want: &queryFinder{packageInfo: &types.Info{}},
		},
		{
			name: "Query",
			qf: &queryFinder{
				packageInfo:    &types.Info{},
				prepInterfaces: map[string]*types.Interface{"querier": &types.Interface{}},
			},
			node: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					Sel: &ast.Ident{Name: "Query"},
				},
				Args: []ast.Expr{&ast.BasicLit{}},
			},
			want: nil,
		},
		{
			name: "QueryContext",
			qf: &queryFinder{
				packageInfo:    &types.Info{},
				prepInterfaces: map[string]*types.Interface{"querierWithContext": &types.Interface{}},
			},
			node: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					Sel: &ast.Ident{Name: "QueryContext"},
				},
				Args: []ast.Expr{&ast.BasicLit{}, &ast.BasicLit{}},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.qf.Visit(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryFinder.Visit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniqueStrings(t *testing.T) {
	tests := []struct {
		name    string
		strings []string
		want    []string
	}{
		{
			name:    "empty slice",
			strings: []string{},
			want:    []string{},
		},
		{
			name:    "all strings are unique",
			strings: []string{"1", "2", "3"},
			want:    []string{"1", "2", "3"},
		},
		{
			name:    "some strings are not unique",
			strings: []string{"1", "2", "3", "2", "1"},
			want:    []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueStrings(tt.strings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniqueStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateCode(t *testing.T) {
	type args struct {
		packageName string
		importPath  string
		queries     []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "format error",
			args: args{
				packageName: "!!!",
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				packageName: "noname",
				importPath:  "github.com/hexdigest/prep",
				queries:     []string{"SELECT", "UPDATE"},
			},
			want: "//go:generate prep -f github.com/hexdigest/prep\n\npackage noname\n\nvar prepStatements = []string{\n\tSELECT,\n\tUPDATE,\n}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateCode(tt.args.packageName, tt.args.importPath, tt.args.queries)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("generateCode() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
