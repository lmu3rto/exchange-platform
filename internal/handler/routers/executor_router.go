package routers

import (
	"github.com/lmu3rto/exchange-platform/internal/handler"
	"net/http"
)

func ExecutorRouter(mux *http.ServeMux, h *handler.ExecutorHandler) {
	mux.HandleFunc("POST /users/{id}/executor", h.Create)
	mux.HandleFunc("GET /users/{id}/executor", h.GetByID)
	mux.HandleFunc("DELETE /users/{id}/executor", h.Delete)
}
