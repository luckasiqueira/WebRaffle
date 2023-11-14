package routes

import (
	"Projetos/rifa/controllers"
	"Projetos/rifa/database"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {

	db := database.Connect()
	look, err := db.Query("select numero, comprador from sorteio order by numero asc")
	if err != nil {
		panic(err.Error())
	}

	// Look for each number in DB and save into slice
	n := controllers.Raffle{}        // each num
	number := []controllers.Raffle{} // slice of instances of nums

	// Iterating DB looking for existing rows
	for look.Next() {

		var num int
		var buyer string

		err := look.Scan(&num, &buyer)
		if err != nil {
			panic(err.Error())
		}
		n.Num, n.Buyer = num, buyer
		number = append(number, n)

	}

	web.ExecuteTemplate(w, "list", number)
	defer db.Close()
}
