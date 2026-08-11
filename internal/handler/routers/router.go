package routers

import (
	"github.com/lmu3rto/exchange-platform/internal/handler"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func NewRouter(user *handler.UserHandler, executor *handler.ExecutorHandler, customer *handler.CustomerHandler) *http.ServeMux {
	mux := http.NewServeMux()

	UserRouter(mux, user)
	ExecutorRouter(mux, executor)
	CustomerRouter(mux, customer)

	mux.Handle("/swagger/{any...}", httpSwagger.WrapHandler)

	return mux
}
