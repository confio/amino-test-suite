package main

import (
	"text/template"
)

var jsTmpl = `{{range .}}const {{.Label}}_binary = "{{.BinaryHex}};"
const {{.Label}}_json = ` + "`{{.JSON}}`;\n\n{{end}}"

func GetJSTemplate() (*template.Template, error) {
	t, err := template.New("js").Parse(jsTmpl)
	return t, err
}
