package handler

import "net/http"

func UserRouter(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /users", h.Create)
}
