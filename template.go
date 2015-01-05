//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

var templateSource = `
package {{.PkgName}}
{{ $variants := .Variants}}

type {{.TypeName}} interface {
	{{range .Variants}} {{.}}() ({{.}}, bool)
	{{end}}
}

type {{.TypeName}}Exhaustive struct {
	{{.TypeName}}

	{{range .Variants}} {{.}}Called bool
	{{end}}
}

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
		{{if eq $recvType . }}
			func (sv {{.}}) {{.}}() ({{.}}, bool) {
				return sv, true
			}
		{{else}}
			func (_ {{$recvType}}) {{.}}() ({{.}}, bool) {
				var v {{.}}
				return v, false
			}
		{{end}}
	{{end}}
{{end}}
`
