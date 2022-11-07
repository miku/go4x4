package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func bar() error {
	if err := foo(); err != nil {
		return fmt.Errorf("bar failed: %w", err)
	}
	return nil
}

func foo() error {
	return fmt.Errorf("foo failed: %w", sql.ErrNoRows)
}

func main() {
	err := bar()
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("data not found, %+v\n", err)
		return
	}
	if err != nil {
		log.Println(err)
	}
}

/* Outputs:
data not found, bar failed: foo failed: sql: no rows in result set
*/
