package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, done FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Printf("%-4s %-30s %s\n", "ID", "TITLE", "DONE")
	fmt.Println("--------------------------------------------")

	for rows.Next() {
		var id int
		var title string
		var done bool
		if err := rows.Scan(&id, &title, &done); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-4d %-30s %v\n", id, title, done)
	}
}
