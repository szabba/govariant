//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

var templateSource = `
package {{.PkgName}}
{{ $typeName := .TypeName }}

// A {{.TypeName}} is one of 
// {{range .Variants}}
//     - {{.}}{{end}}
//
// In each implementation exactly one of the methods should have the second
// return value be true.
type {{.TypeName}} interface {
	{{range .Variants}}
	// {{.}} returns a {{.}} and a boolean. When the second return value is true,
	// the {{$typeName}} is the returned {{.}}.
	{{.}}() ({{.}}, bool)
	{{end}}
}

// A core{{.TypeName}} provides default implementations for the methods of a
// Shape. The wrapper types only redefine the one for the associated variant.
type core{{.TypeName}} struct{}

{{range .Variants}}
func (_ core{{$typeName}}) {{.}}() ({{.}}, bool) {
	var v {{.}}
	return v, false
}
{{end}}

{{range .Variants}}
// {{$typeName}} converts a {{.}} to an instance of the sum type {{$typeName}}.
func (v {{.}}) {{$typeName}}() {{$typeName}} {
	return wrap{{.}}{{$typeName}}{
		wrapped{{.}}: v,
	}
}

type wrap{{.}}{{$typeName}} struct {
	core{{$typeName}}
	wrapped{{.}} {{.}}
}

func (w wrap{{.}}{{$typeName}}) {{.}}() ({{.}}, bool) {
	return w.wrapped{{.}}, true
}
{{end}}

{{if .ExhaustionChecker}}
// A {{.TypeName}}Exhaustive is a {{.TypeName}} that can be used to check
// exhaustivity in tests
type {{.TypeName}}Exhaustive struct {
	{{.TypeName}}

	{{range .Variants}} {{.}}Called bool
	{{end}}
}

// Checks whether all the variants of the {{.TypeName}} were considered.
func (se {{.TypeName}}Exhaustive) Exhaustive() bool {
	{{range .Variants}}
	if !se.{{.}}Called {
		return false
	}
	{{end}}

	return true
}
{{end}}
`
