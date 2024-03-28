package main

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"
)

//go:embed sample.tmpl
var SampleTmpl string

type SampleTmplData struct {
	Name string
	Age  int
}

func main() {
	data := SampleTmplData{
		Name: "name",
		Age:  10,
	}

	tmpl, err := template.New("").Parse(SampleTmpl)
	if err != nil {
		panic(err)
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
