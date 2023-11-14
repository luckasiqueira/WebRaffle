package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Connect() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbdata := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", user, dbname, pass, host, port)

	info := dbdata
	db, err := sql.Open("postgres", info)
	if err != nil {
		panic(err.Error())
	}
	return db
}
