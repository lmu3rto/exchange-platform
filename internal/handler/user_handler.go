package handler

import (
	"github.com/lmu3rto/exchange-platform/internal/domain/models"
	"net/http"
	"encoding/json"

)


func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userService.Create(ctx, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createdUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}