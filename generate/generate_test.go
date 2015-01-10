package generate

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/kr/pretty"
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

	mustBeInterface(t, f, sumType)
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

	typ := mustBeInterface(t, f, sumType)

	namesFound := make(map[string]bool)
	for _, method := range typ.Methods.List {
		name := method.Names[0].Name
		namesFound[name] = true
	}

	for _, variant := range variants {
		if !namesFound[variant] {
			t.Errorf("generated type should contain method named %s", variant)
		}
	}
}

func TestGeneratedSumTypesHaveMethodsWithExpectedResultTypes(t *testing.T) {
	pkg := "pkg"
	sumType := "Sum"
	variants := []string{"X", "Y", "Z"}

	src := Generate(pkg, sumType, variants...)

	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "src.go", src, 0)

	typ := mustBeInterface(t, f, sumType)

	variantSet := stringSet(variants...)
	for _, method := range typ.Methods.List {

		name := method.Names[0].Name
		if variantSet[name] {

			funcTyp, isFunc := method.Type.(*ast.FuncType)
			if isFunc {

				hasResults(t, sumType, name, funcTyp, name, "bool")
			}
		}
	}
}

// mustBeInterface ensures that type f contains a type named typName and that
// it's an interface, returning it's AST representation. Otherwise it
// immediately fails the test.
func mustBeInterface(t *testing.T, f *ast.File, typName string) *ast.InterfaceType {
	typ, ok := typeNamed(f, typName)
	if !ok {
		t.Fatalf("generated source must contain type declaration for %s type", typName)
	}

	asInterface, ok := typ.Type.(*ast.InterfaceType)
	if !ok {
		t.Fatalf("generated %s type should be an interface not a %T", typName, typ.Type)
	}

	return asInterface
}

var _ = fmt.Println
var _ = pretty.Println

// hasResults checks that the named method of the named sumType has the
// specified type names by inspecting it's signature sig.
func hasResults(t *testing.T, sumType, funcName string, sig *ast.FuncType, typenames ...string) {
	results := resultTypes(sig)

	if len(results) != len(typenames) {
		t.Errorf("method %s of type %s should have two return values", funcName, sumType)
		return
	}

	for i, actualResultType := range results {
		expected := typenames[i]

		if actualResultType != expected {
			t.Errorf("%d-th result type of %s's method %s should be %s, not %s",
				i, sumType, funcName, expected, actualResultType)
		}
	}
}

func resultTypes(sig *ast.FuncType) []string {
	types := []string{}

	for _, group := range sig.Results.List {
		typ := group.Type.(*ast.Ident).Name

		for i := 0; i < len(group.Names); i++ {
			types = append(types, typ)
		}
		if len(group.Names) == 0 {
			types = append(types, typ)
		}
	}

	return types
}

func stringSet(elems ...string) map[string]bool {
	set := make(map[string]bool)
	for _, elem := range elems {
		set[elem] = true
	}
	return set
}
