package main

import (
	"text/template"
)

var jsTmpl = `
const {{.Label}}_json = ` +
	"`{{.JSON}}`;\n" +
	`const {{.Label}}_binary = "{{.BinaryHex}};"
	`

func GetJSTemplate() (*template.Template, error) {
	t, err := template.New("js").Parse(jsTmpl)
	return t, err
}
