package generate

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestGeneratedSourceParsesWithoutErrors(t *testing.T) {
	pkg := "pkg"
	sumType := "Sum"
	variants := []string{"X", "Y", "Z"}

	src := Generate(pkg, sumType, variants...)

	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		t.Error(err)
	}
}

func TestGeneratedSourceDeclaresTypeWithNameGiven(t *testing.T) {
	pkg := "pkg"
	sumType := "Sum"
	variants := []string{"X", "Y", "Z"}

	src := Generate(pkg, sumType, variants...)

	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "src.go", src, 0)

	found := false
	for _, typ := range typesInFile(f) {
		if typ.Name.Name == sumType {
			found = true
		}
	}

	if !found {
		t.Errorf("type %s not declared in generated sources", sumType)
	}
}

func typesInFile(f *ast.File) []*ast.TypeSpec {
	var types []*ast.TypeSpec

	for _, decl := range f.Decls {
		decl := decl.(*ast.GenDecl)

		if decl.Tok == token.TYPE {

			for _, spec := range decl.Specs {
				if typ, ok := spec.(*ast.TypeSpec); ok {

					types = append(types, typ)
				}
			}
		}
	}

	return types
}

func TestGeneratedSourceHasRightPackageClause(t *testing.T) {
	for _, pkg := range []string{"pkg", "pack"} {
		sumType := "Sum"
		variants := []string{"X", "Y", "Z"}

		src := Generate(pkg, sumType, variants...)

		fset := token.NewFileSet()
		f, _ := parser.ParseFile(fset, "src.go", src, parser.PackageClauseOnly)

		if f.Name.Name != pkg {
			t.Errorf("generated file should belong to package %s not %s",
				pkg, f.Name.Name)
		}
	}
}

func TestGeneratedSumTypeIsInterface(t *testing.T) {
	pkg := "pkg"
	sumType := "Sum"
	variants := []string{"X", "Y", "Z"}

	src := Generate(pkg, sumType, variants...)

	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "src.go", src, 0)

	typ, ok := typeNamed(f, sumType)
	if !ok {
		t.Fatalf("generated source must contain type declaration for %s type", sumType)
	}

	_, isInterface := typ.Type.(*ast.InterfaceType)
	if !isInterface {
		t.Errorf("generated %s type should be an interface not a %T", sumType, typ.Type)
	}
}

// Looks for the declaration of a type named name in the specified file. The
// second return value is false if a type with the given name is not declared
// in the file.
func typeNamed(f *ast.File, name string) (*ast.TypeSpec, bool) {
	for _, typ := range typesInFile(f) {
		if typ.Name.Name == name {
			return typ, true
		}
	}
	return nil, false
}

func TestGeneratedSumTypeHasMethodsWithExpectedNames(t *testing.T) {
	pkg := "pkg"
	sumType := "Sum"
	variants := []string{"X", "Y", "Z"}

	src := Generate(pkg, sumType, variants...)

	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "src.go", src, 0)

	typ, ok := typeNamed(f, sumType)
	if !ok {
		t.Fatalf("generated source must contain type declaration for %s type", sumType)
	}

	asInterface, ok := typ.Type.(*ast.InterfaceType)
	if !ok {
		t.Errorf("generated %s type should be an interface not a %T", sumType, typ.Type)
	}

	namesFound := make(map[string]bool)
	for _, method := range asInterface.Methods.List {
		name := method.Names[0].Name
		namesFound[name] = true
	}

	for _, variant := range variants {
		if !namesFound[variant] {
			t.Errorf("generated type should contain method named %s", variant)
		}
	}
}
