package controllers

import (
	"Projetos/rifa/database"
	"time"
)

type Complete struct {
	Status bool
	When   string
	Num    string
	Buyer  string
}

var complete = &Complete{}

func Define() *Complete {
	db := database.Connect()
	getstatus, err := db.Query("select done from status where id = 1;")
	if err != nil {
		panic(err.Error())
	}

	var status string
	for getstatus.Next() {
		getstatus.Scan(&status)
	}

	if status == "true" {
		complete.Status = true
		complete = Save()
	} else {
		complete.Status = false
	}

	db.Close()
	return complete
}

func Run() {
	db := database.Connect()
	defer db.Close()

	find, err := db.Query("select numero, comprador from sorteio order by random() limit 1;")
	if err != nil {
		panic(err.Error())
	}

	var a int
	var b string
	d := time.Now()
	date := d.Format("15:04:05 de 02-01-2006")

	for find.Next() {
		find.Scan(&a, &b)
	}

	save, err := db.Prepare("UPDATE public.status SET done = true, winner = $1, moment = $2, number = $3 WHERE id = 1;")
	save.Exec(b, date, a)
}

func Save() *Complete {
	db := database.Connect()

	winner, err := db.Query("select winner, number, moment from status where id = 1;")
	if err != nil {
		panic(err.Error())
	}

	for winner.Next() {
		var a, b, c string
		winner.Scan(&a, &b, &c)
		complete.Buyer = a
		complete.Num = b
		complete.When = c
	}

	complete.Status = true
	db.Close()
	return complete
}
