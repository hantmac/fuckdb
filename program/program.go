package program

import (
	"errors"
	"go/ast"
	"go/parser"

	"github.com/fatih/astrewrite"

	"golang.org/x/tools/go/loader"
)

type Program struct {
	*loader.Program
}

// var (
// 	ErrPkgNotFound    = errors.New("pkg not found")
// 	ErrStructNotFound = errors.New("struct not found")
// )

func NewProgram(pkgs []string) (*Program, error) {
	conf := loader.Config{}
	conf.ParserMode |= parser.ParseComments
	for _, pkg := range pkgs {
		conf.Import(pkg)
	}
	prog, err := conf.Load()
	if err != nil {
		return nil, err
	}

	return &Program{
		Program: prog,
	}, nil

}

func (p *Program) GetPkgByName(name string) (*loader.PackageInfo, error) {
	if name == "" {
		return nil, errors.New("empty package name")
	}
	for _, pkg := range p.InitialPackages() {
		if pkg.String() == name {
			return pkg, nil
		}
	}

	return nil, errors.New("pkg " + name + " not found")
}

func (p *Program) GetStructByName(pkgName, structName string) (typeSpec *ast.TypeSpec, err error) {
	pi, err := p.GetPkgByName(pkgName)
	if err != nil {
		return nil, err
	}

	// var typeSpec *ast.TypeSpec
	// var err error
	for _, file := range pi.Files {
		// astrewrite.Walk(file, func(node ast.Node) (ast.Node, bool) {
		// 	ts, ok := node.(*ast.TypeSpec)
		// 	if ok && ts.Name.Name == structName {
		// 		typeSpec = ts
		// 		return node, false
		// 	}
		// 	return node, true
		// })
		typeSpec, err = GetStructByName(file, structName)
		if err == nil {
			// return typeSpec, nil
			return
		}
	}
	// if typeSpec != nil {
	// 	return typeSpec, nil
	// }

	// return nil, errors.New("type " + structName + " not found")
	return
}

func (p *Program) GetStructByNameFromPkgs(pkgs []string, structName string) (*ast.TypeSpec, error) {
	for _, pkg := range pkgs {
		if typeSpec, err := p.GetStructByName(pkg, structName); err == nil {
			return typeSpec, nil
		}
	}

	return nil, errors.New("type " + structName + " not found")
}

// ExtractStruct extract struct spec to all type spec it depends (including itself)
func (p *Program) ExtractStruct(node ast.Node, sub *[]*ast.TypeSpec) {

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
		if !p.containType(t.Name.Name, *sub) {
			*sub = append(*sub, t)
			p.ExtractStruct(t.Type, sub)
		}
	case *ast.StructType:
		for _, field := range t.Fields.List {
			p.ExtractStruct(field.Type, sub)
		}
	case *ast.StarExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			p.ExtractStruct(ident, sub)
		}
	case *ast.MapType:
		p.ExtractStruct(t.Value, sub)
	case *ast.Ident:
		if t.Obj == nil {
			return
		}
		typeSpec, ok := t.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			return
		}
		if p.containType(typeSpec.Name.Name, *sub) {
			return
		}
		*sub = append(*sub, typeSpec)
		p.ExtractStruct(typeSpec.Type, sub)
	}
}

func (p *Program) containType(name string, types []*ast.TypeSpec) bool {
	for _, typ := range types {
		if name == typ.Name.Name {
			return true
		}
	}
	return false
}

func (p *Program) GetFuncByName(pkgName, funcName string) (*ast.FuncDecl, error) {
	pi, err := p.GetPkgByName(pkgName)
	if err != nil {
		return nil, err
	}
	var funcDecl *ast.FuncDecl

	for _, file := range pi.Files {
		astrewrite.Walk(file, func(node ast.Node) (ast.Node, bool) {
			fd, ok := node.(*ast.FuncDecl)
			if !ok {
				return node, true
			}

			if fd.Name.Name == funcName {
				funcDecl = fd
				return node, false
			}
			return node, true
		})

	}

	if funcDecl != nil {
		return funcDecl, nil
	}

	return nil, errors.New("func " + funcName + " not found")
}

func (p *Program) GetValueByName(pkgName, valueName string) (*ast.ValueSpec, error) {
	pi, err := p.GetPkgByName(pkgName)
	if err != nil {
		return nil, err
	}
	var valueSpec *ast.ValueSpec

	for _, file := range pi.Files {
		astrewrite.Walk(file, func(node ast.Node) (ast.Node, bool) {
			spec, ok := node.(*ast.ValueSpec)
			if !ok {
				return node, true
			}

			if spec.Names[0].Name == valueName {
				valueSpec = spec
				return node, false
			}
			return node, true
		})

	}

	if valueSpec != nil {
		return valueSpec, nil
	}

	return nil, errors.New("value " + valueName + " not found")
}
