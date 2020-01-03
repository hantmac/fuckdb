package program

import (
	"errors"
	"go/ast"

	"github.com/fatih/astrewrite"
)

func GetStructByName(f *ast.File, structName string) (*ast.TypeSpec, error) {
	var typeSpec *ast.TypeSpec
	astrewrite.Walk(f, func(node ast.Node) (ast.Node, bool) {
		ts, ok := node.(*ast.TypeSpec)
		if ok && ts.Name.Name == structName {
			typeSpec = ts
			return node, false
		}
		return node, true
	})

	if typeSpec != nil {
		return typeSpec, nil
	}

	return nil, errors.New("type " + structName + " not found")
}

type FuncMatchStruct func(structName string) bool

func FindMatchStruct(files []*ast.File, match FuncMatchStruct) []*ast.TypeSpec {
	if match == nil {
		return nil
	}

	types := []*ast.TypeSpec{}
	for _, f := range files {
		astrewrite.Walk(f, func(node ast.Node) (ast.Node, bool) {
			ts, ok := node.(*ast.TypeSpec)
			if ok {
				if _, isStruct := ts.Type.(*ast.StructType); !isStruct {
					return node, true
				}
				if ts.Name != nil && match(ts.Name.Name) && !ContainType(ts.Name.Name, types) {
					types = append(types, ts)
				}
				return node, true
			}
			return node, true
		})

	}

	return types
}

// ExtractStruct extract struct spec to all type spec it depends (including itself)
func ExtractStruct(node ast.Node, sub *[]*ast.TypeSpec) {

	// type BarReq struct {
	// 	A int `json:"a"`
	// 	B struct {						// struct type
	// 		SubA int `json:"sub_a"`
	// 	} `json:"b"`
	// 	C TC                `json:"c"`	// ident type
	// 	D models.Group      `json:"d"`	// select expr
	// 	E map[string]string `json:"e"`	// map type
	//	G *TG							// star expr
	// }

	switch t := node.(type) {
	case *ast.TypeSpec:
		if !ContainType(t.Name.Name, *sub) {
			*sub = append(*sub, t)
			ExtractStruct(t.Type, sub)
		}
	case *ast.StructType:
		for _, field := range t.Fields.List {
			ExtractStruct(field.Type, sub)
		}
	case *ast.StarExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			ExtractStruct(ident, sub)
		}
	case *ast.MapType:
		ExtractStruct(t.Value, sub)
	case *ast.Ident:
		if t.Obj == nil {
			return
		}
		typeSpec, ok := t.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			return
		}
		if ContainType(typeSpec.Name.Name, *sub) {
			return
		}
		*sub = append(*sub, typeSpec)
		ExtractStruct(typeSpec.Type, sub)
	}
}

func ContainType(name string, types []*ast.TypeSpec) bool {
	for _, typ := range types {
		if name == typ.Name.Name {
			return true
		}
	}
	return false
}
