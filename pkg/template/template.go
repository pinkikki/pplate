package template

import (
	"text/template"
)

type Template struct {
	*template.Template
}

func CreateTemplate(name string, content string) *Template {
	return &Template{template.Must(template.New(name).Parse(content))}
}
