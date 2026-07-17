package main

import (
	"log"

	"fmt"
	"github.com/lmu3rto/exchange-platform/internal/database"
)

const databaseURL string = "postgres://mu3rto:platform_password@localhost:5432/exchange_platform?sslmode=disable"

func main() {
	db, err := database.New(databaseURL)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Database connected!")
}
