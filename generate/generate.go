package generate

// Generate returns the source for a file that defines a sum type when given
// it's name and the names of it's variants.
func Generate(sumType string, variants ...string) string {
	return "package pkg"
}
