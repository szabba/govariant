package generate

import "fmt"

var pkgFormat = `
package %s

type %s interface{}
`

// Generate returns the source for a file that defines a sum type when given
// it's name and the names of it's variants.
func Generate(pkg, sumType string, variants ...string) string {
	return fmt.Sprintf(pkgFormat, pkg, sumType)
}
