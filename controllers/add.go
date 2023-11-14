package controllers

import (
	"Projetos/rifa/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Raffle struct {
	Buyer   string
	Num     int
	Contact string
}

var confirm string

func Add(w http.ResponseWriter, r *http.Request) {
	contact := r.FormValue("contato")
	buyer := strings.ToUpper(r.FormValue("comprador"))
	n := r.FormValue("numero")
	num, err := strconv.Atoi(n)

	done := false
	if err != nil {
		fmt.Print("Error when converting num. \nError:", err)
	}

	// Add into db
	db := database.Connect()
	insert, err := db.Prepare("insert into sorteio (numero, comprador, contato) values($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(num, buyer, contact)

	done = true
	if done {
		http.Redirect(w, r, "/lista", 301)
	}
	defer db.Close()
}
