package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lmu3rto/exchange-platform/internal/database"
	"github.com/lmu3rto/exchange-platform/internal/handler"
	"github.com/lmu3rto/exchange-platform/internal/repository"
	"github.com/lmu3rto/exchange-platform/internal/service"
)

var (
	ReadTimeout  = 5 * time.Second
	WriteTimeout = 15 * time.Second
	IdleTimeout  = 120 * time.Second
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := database.New(databaseURL)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	h := handler.NewHandler(userService)

	mux := http.NewServeMux()

	handler.NewRouter(h)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		IdleTimeout:  IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Printf("server failed: %v", err)
		return
	}

}
