package routers

import (
	"github.com/lmu3rto/exchange-platform/internal/handler"
	"net/http"
)

func CustomerRouter(mux *http.ServeMux, h *handler.CustomerHandler) {
	mux.HandleFunc("POST /users/{id}/customer", h.Create)
	mux.HandleFunc("GET /users/{id}/customer", h.GetByID)
	mux.HandleFunc("DELETE /users/{id}/customer", h.Delete)
}
