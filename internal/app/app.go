package app

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmu3rto/exchange-platform/internal/config"
	"github.com/lmu3rto/exchange-platform/internal/database"
	"github.com/lmu3rto/exchange-platform/internal/handler"
	"github.com/lmu3rto/exchange-platform/internal/handler/routers"
	"github.com/lmu3rto/exchange-platform/internal/repository"
	"github.com/lmu3rto/exchange-platform/internal/service"
)

type App struct {
	DB      *pgxpool.Pool
	Handler http.Handler
}

func New(cfg config.Config, logger *slog.Logger) (*App, error) {
	db, err := database.New(cfg.DatabaseURL)

	if err != nil {
		return nil, fmt.Errorf("create database: %w", err)
	}

	userRepo := repository.NewUserRepository(db)
	execRepo := repository.NewExecutorRepository(db)
	custRepo := repository.NewCustomerRepository(db)

	userService := service.NewUserService(userRepo)
	execService := service.NewExecutorService(execRepo, userRepo)
	custService := service.NewCustomerService(custRepo, userRepo)

	user := handler.NewUserHandler(*userService, logger)
	exec := handler.NewExecutorHandler(*execService, logger)
	cust := handler.NewCustomerHandler(*custService, logger)

	mux := routers.NewRouter(user, exec, cust)

	return &App{
		DB:      db,
		Handler: mux,
	}, nil
}
