package handler

import "net/http"

func NewRouter(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()

	UserRouter(mux, h)

	return mux
}
