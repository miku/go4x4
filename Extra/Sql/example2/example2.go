package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var (
		id   int
		name string
	)

	rows, err := db.Query("select id, name from entries where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // will be called implicitly, when rows.Next returns false, but still ok to use; ok to close closed rows

	for rows.Next() { // similar to bufio.Scanner, may stop and end or error
		err := rows.Scan(&id, &name) // https://pkg.go.dev/database/sql#Rows.Scan -- performs various ops implicitly

		// Scan also converts between string and numeric types, as long as no information would be lost.

		// Can read raw data into byte slice, or can return RawBytes to avoid copy.

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}

	err = rows.Err() // always check, whether we got an error
	if err != nil {
		log.Fatal(err)
	}
}
