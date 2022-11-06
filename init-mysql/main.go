package main

import (
	"database/sql"
	"log"
)

func main() {

	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	log.Print(db)
}
