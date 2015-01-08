package generate

import "fmt"

// Generate returns the source for a file that defines a sum type when given
// it's name and the names of it's variants.
func Generate(pkg, sumType string, variants ...string) string {
	out := fmt.Sprintf("package %s\n\n", pkg)
	out += fmt.Sprintf("type %s interface {\n", sumType)
	for _, variant := range variants {
		out += fmt.Sprintf("\t%s() (bool,  int)\n", variant)
	}
	out += "}"
	return out
}
