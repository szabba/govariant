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
