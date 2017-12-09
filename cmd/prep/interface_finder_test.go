package main

import (
	"go/ast"
	"go/types"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/loader"
)

func Test_findInterfaces(t *testing.T) {
	conf := loader.Config{
		TypeChecker: types.Config{
			FakeImportC:              true,
			DisableUnusedImportCheck: true,
		},
		TypeCheckFuncBodies: func(path string) bool {
			return false
		},
	}

	conf.Import(importPath)

	prog, err := conf.Load()
	require.NoError(t, err)

	pkg := prog.Package(importPath)
	require.NotNil(t, pkg)

	interfaces := findInterfaces(pkg)
	assert.Len(t, interfaces, 12)
}

func Test_interfaceFinder_Visit(t *testing.T) {
	tests := []struct {
		name string
		i    *interfaceFinder
		node ast.Node
		want ast.Visitor
	}{
		{
			name: "node is not a type spec",
			i:    &interfaceFinder{},
			want: &interfaceFinder{},
		},
		{
			name: "type spec is not an interface",
			i:    &interfaceFinder{packageInfo: &types.Info{}},
			node: &ast.TypeSpec{
				Name: &ast.Ident{Name: "noname"},
				Type: &ast.Ident{Name: "sometype"},
			},
			want: nil,
		},
		{
			name: "type spec is not an interface",
			i:    &interfaceFinder{packageInfo: &types.Info{}},
			node: &ast.TypeSpec{
				Name: &ast.Ident{Name: "noname"},
				Type: &ast.Ident{Name: "sometype"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Visit(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("interfaceFinder.Visit() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("type spec is an interface", func(t *testing.T) {
		node := &ast.TypeSpec{
			Name: &ast.Ident{Name: "noname"},
			Type: &ast.Ident{Name: "sometype"},
		}

		i := &interfaceFinder{
			packageInfo: &types.Info{
				Types: map[ast.Expr]types.TypeAndValue{
					node.Type: {Type: &types.Interface{}},
				},
			},
			foundInterfaces: map[string]*types.Interface{},
		}

		res := i.Visit(node)
		assert.Nil(t, res)
		assert.NotNil(t, i.foundInterfaces["noname"])
	})
}
