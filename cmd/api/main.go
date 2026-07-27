package main

import (
	"log"
	"net/http"

	"github.com/lmu3rto/exchange-platform/internal/database"
	"github.com/lmu3rto/exchange-platform/internal/handler"
	"github.com/lmu3rto/exchange-platform/internal/repository"
	"github.com/lmu3rto/exchange-platform/internal/service"
)

const databaseURL string = "postgres://mu3rto:platform_password@localhost:5432/exchange_platform?sslmode=disable"

func main() {
	db, err := database.New(databaseURL)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()


	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)


	h := handler.NewHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", h.Create)

	http.ListenAndServe(":8080", mux)

}
