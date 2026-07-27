package handler

import (
	"github.com/lmu3rto/exchange-platform/internal/domain/models"
	"net/http"
	// "encoding/json"

)


func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user models.User

	if err := decodeJSON(*r, &user); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	createdUser, err := h.userService.Create(ctx, user)

	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
	}


	if err := writeJSON(w, http.StatusCreated, createdUser); err != nil {
		writeError(w, err, http.StatusInternalServerError)
	}
}