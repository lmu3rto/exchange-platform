package handler

import "net/http"

func UserRouter(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /users", h.Create)
	mux.HandleFunc("GET /users", h.GetByName)
	mux.HandleFunc("GET /users/{id}", h.GetByID)
	mux.HandleFunc("PATCH /users/{id}", h.Update)
	mux.HandleFunc("PUT /users/{id}", h.Update)
}
