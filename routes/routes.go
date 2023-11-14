package routes

import (
	"net/http"
	"text/template"
)

// Import .html files as templates
var web = template.Must(template.ParseGlob("public/*html"))

func Index(w http.ResponseWriter, r *http.Request) {
	web.ExecuteTemplate(w, "index", nil)

}
