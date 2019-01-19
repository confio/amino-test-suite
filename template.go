package main

import (
	"text/template"
)

var jsTmpl = `{{range .}}const {{.Label}}_binary = '{{.BinaryHex}};'
{{if .SignBytes}}const {{.Label}}_sign = '{{.SignBytes}}'
{{end}}const {{.Label}}_json = {{.JSON}}

{{end}}`

func GetJSTemplate() (*template.Template, error) {
	t, err := template.New("js").Parse(jsTmpl)
	return t, err
}
