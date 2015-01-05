//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

var config struct {
	PkgName  string
	TypeName string
	Variants []string
}

var usage = `
govariant [OPTIONS...] SUM_TYPE VARIANT_ONE VARIANT_TWO [OTHER_VARIANTS...]

SUM_TYPE is the name of the generated type. The VARIANTS are the variant types;
there must be at least two of them.
`[1:]

func init() {

	flag.StringVar(&config.PkgName, "pkg", os.Getenv("GOPACKAGE"),
		"specifies the package name; defaults to $GOPACKAGE")

	flag.Parse()

	switch len(flag.Args()) {
	case 0:
		exitWithUsage("result type name not specified")
	case 1:
		exitWithUsage("no variants specified")
	case 2:
		exitWithUsage("only one variant specified")
	}

	if config.PkgName == "" {
		exitWithUsage(
			"set the package either throug the -pkg flag or GOPACKAGE " +
				"environment variable")
	}

	config.TypeName = flag.Args()[0]
	config.Variants = flag.Args()[1:]
}

func main() {
	t := template.Must(template.New("variants").Parse(templateSource))

	var b bytes.Buffer
	err := t.Execute(&b, config)
	if err != nil {
		log.Fatal(err)
	}

	fmted, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(config.TypeName+"_variant.go", fmted, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func exitWithUsage(msg string) {
	fmt.Printf("%#v\n", flag.Args())
	fmt.Println(msg)
	fmt.Print(usage)
	os.Exit(1)
}
