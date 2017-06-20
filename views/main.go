package views

import (
	"html/template"
	"net/http"
)

// T ...
var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("./views/*.html"))
}

// Render ...
func Render(w http.ResponseWriter, name string, data interface{}) (err error) {
	if err = t.ExecuteTemplate(w, name, data); err != nil {
		return
	}
	return nil
}
