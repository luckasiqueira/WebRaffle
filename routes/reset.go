package routes

import (
	"Projetos/rifa/controllers"
	"net/http"
)

func Reset(w http.ResponseWriter, r *http.Request) {
	controllers.Clear(w, r)
	web.ExecuteTemplate(w, "clear", nil)
}
