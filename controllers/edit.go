package controllers

import (
	"Projetos/rifa/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Update(w http.ResponseWriter, r *http.Request) {
	done := false
	i := r.FormValue("alterar")
	buyer := strings.ToUpper(r.FormValue("comprador"))
	n := r.FormValue("numero")
	contact := r.FormValue("contato")
	num, err := strconv.Atoi(n)
	id, err := strconv.Atoi(i)

	if err != nil {
		fmt.Print("Error when converting num. \nError:", err)
	}

	db := database.Connect()
	query := fmt.Sprintf("UPDATE public.sorteio SET numero = '%d', comprador = '%s', contato = '%s' WHERE numero = %d;", num, buyer, contact, id)
	//update, err := db.Prepare("UPDATE public.sorteio SET numero = '$1', comprador = '$2', contato = '$3' WHERE numero = $4;")

	update, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	update.Exec()
	done = true
	if done {
		http.Redirect(w, r, "/lista", 301)
	}
	defer db.Close()
}
