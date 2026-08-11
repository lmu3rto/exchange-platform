package main

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/lmu3rto/exchange-platform/docs"
	"github.com/lmu3rto/exchange-platform/internal/app"
	"github.com/lmu3rto/exchange-platform/internal/config"
)

var (
	ReadTimeout  = 5 * time.Second
	WriteTimeout = 15 * time.Second
	IdleTimeout  = 120 * time.Second
)

// @title Exchange Platform API
// @version 1.0
// @description API биржи фриланса.
// @host localhost:8080
// @BasePath /
func main() {

	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux, err := app.New(cfg, logger)

	if err != nil {
		logger.Error("failed to connect db", "error", err)
		return
	}

	defer mux.DB.Close()

	srv := &http.Server{
		Addr:         cfg.HttpAddr,
		Handler:      mux.Handler,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		IdleTimeout:  IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Printf("server failed: %v", err)
		return
	}

}
