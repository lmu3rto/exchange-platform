package routers

import (
	"github.com/lmu3rto/exchange-platform/internal/handler"
	"net/http"
)

func UserRouter(mux *http.ServeMux, h *handler.UserHandler) {
	mux.HandleFunc("POST /users", h.Create)
	mux.HandleFunc("GET /users/name", h.GetByName)
	mux.HandleFunc("GET /users/{id}", h.GetByID)
	mux.HandleFunc("PATCH /users/{id}", h.Update)
	mux.HandleFunc("PUT /users/{id}", h.Update)
}
