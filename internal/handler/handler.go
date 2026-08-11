package handler

import (
	"log/slog"

	"github.com/lmu3rto/exchange-platform/internal/service"
)

type UserHandler struct {
	userService service.UserService
	logger      *slog.Logger
}

func NewUserHandler(us service.UserService, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		userService: us,
		logger:      logger,
	}
}

type ExecutorHandler struct {
	executorService service.ExecutorService
	logger          *slog.Logger
}

func NewExecutorHandler(es service.ExecutorService, logger *slog.Logger) *ExecutorHandler {
	return &ExecutorHandler{
		executorService: es,
		logger:          logger,
	}
}

type CustomerHandler struct {
	customerService service.CustomerService
	logger          *slog.Logger
}

func NewCustomerHandler(cs service.CustomerService, logger *slog.Logger) *CustomerHandler {
	return &CustomerHandler{
		customerService: cs,
		logger:          logger,
	}
}
