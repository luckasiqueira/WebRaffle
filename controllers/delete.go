package controllers

import (
	"Projetos/rifa/database"
	"fmt"
	"net/http"
	"strconv"
)

func Remove(w http.ResponseWriter, r *http.Request) {
	done := false
	i := r.FormValue("alterar")
	id, err := strconv.Atoi(i)

	if err != nil {
		fmt.Print("Error when converting num. \nError:", err)
	}

	db := database.Connect()
	query := fmt.Sprintf("DELETE	FROM public.sorteio	WHERE numero = %d", id)
	fmt.Printf(query)
	remove, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	remove.Exec()
	done = true
	if done {
		http.Redirect(w, r, "/lista", 301)
	}
	defer db.Close()
}
