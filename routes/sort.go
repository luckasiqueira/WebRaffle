package routes

import (
	"Projetos/rifa/controllers"

	"net/http"
)

func Sort(w http.ResponseWriter, r *http.Request) {
	complete := controllers.Define()
	web.ExecuteTemplate(w, "sort", complete)
}

func Draw(w http.ResponseWriter, r *http.Request) {
	controllers.Run()
	complete := controllers.Define()
	web.ExecuteTemplate(w, "draw", complete)
}
