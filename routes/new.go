package routes

import (
	"net/http"
)

func New(w http.ResponseWriter, r *http.Request) {
	web.ExecuteTemplate(w, "new", nil)
}
