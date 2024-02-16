package template

import "html/template"

var Templates = template.Must(template.ParseGlob("internal/template/*.html"))
