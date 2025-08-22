package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// UPDATE de dados
	result, err := db.Exec("UPDATE users SET name = ? WHERE id = ?", "James Bond", 3)
	if err != nil {
		log.Fatal(err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Updated %d row(s)\n", affectedRows)

	// DELETE de dados
	result, err = db.Exec("DELETE FROM users WHERE id = ?", 2)
	if err != nil {
		log.Fatal(err)
	}
	affectedRows, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted %d row(s)\n", affectedRows)
}
