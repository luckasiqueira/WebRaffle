package controllers

import (
	"Projetos/rifa/database"
	"net/http"
)

func Clear(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()

	cleanWinner, err := db.Prepare("UPDATE public.status SET done = null WHERE id = 1;")
	if err != nil {
		panic(err.Error())
	}

	cleanWinner.Exec()

	http.Redirect(w, r, "/sorteio", 301)
	defer db.Close()
}
