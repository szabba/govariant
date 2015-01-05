//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

var templateSource = `
package {{.PkgName}}
{{ $typeName := .TypeName }}
{{ $variants := .Variants }}

// A {{.TypeName}} is one of 
// {{range .Variants}}
//     - {{.}}{{end}}
type {{.TypeName}} interface {
	{{range .Variants}} {{.}}() ({{.}}, bool)
	{{end}}
}

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

{{range $recvType := .Variants}}
	{{range $variants}}
		// {{.}} implements the corresponding method of {{$typeName}} on the {{$recvType}} type.
		{{if eq $recvType . }} func (sv {{.}}) {{.}}() ({{.}}, bool) {
				return sv, true
			}
		{{else}} func (_ {{$recvType}}) {{.}}() ({{.}}, bool) {
				var v {{.}}
				return v, false
			}
		{{end}}
	{{end}}
{{end}}
`
