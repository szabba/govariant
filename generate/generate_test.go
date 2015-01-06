package generate

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestGeneratedSourceParsesWithoutErrors(t *testing.T) {
	sumType := "Sum"
	variants := []string{"X", "Y", "Z"}

	src := Generate(sumType, variants...)

	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		t.Error(err)
	}
}
