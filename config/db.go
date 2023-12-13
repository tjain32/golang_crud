package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	//DB, _ = sql.Open("postgres", "postgres://<username>:<pwd>@localhost:5432/inventory?sslmode=disable")
	//create the table if it doesn't exist
	_, err1 := DB.Exec("CREATE TABLE IF NOT EXISTS books (isbn TEXT PRIMARY KEY, title TEXT, author TEXT, price BIGINT)")
	if err1 != nil {
		fmt.Println(err1)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("You connected to your database.")
}
