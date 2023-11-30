package main

import (
	"Projetos/rifa/controllers"
	"Projetos/rifa/routes"
	"net/http"
)

// Listen and serve
func main() {
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/sorteio", routes.Sort)
	http.HandleFunc("/rodar", routes.Draw)
	http.HandleFunc("/lista", routes.List)
	http.HandleFunc("/cadastro", routes.New)
	http.HandleFunc("/add", controllers.Add)
	http.HandleFunc("/editar", routes.Edit)
	http.HandleFunc("/remover", routes.Delete)
	http.HandleFunc("/resetar", routes.Reset)

	// File server
	static := http.FileServer(http.Dir("public/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	http.ListenAndServe(":8881", nil)
}
