package handler

import (
	"errors"
	"log"
	"net/http"

	// "encoding/json"
	"github.com/lmu3rto/exchange-platform/internal/domain/models"
	"github.com/lmu3rto/exchange-platform/internal/handler/dto"
	"github.com/lmu3rto/exchange-platform/internal/service"
	"strings"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

var (
	ErrNameLong  = errors.New("Name is too long")
	ErrNameShort = errors.New("Name is too short")
	ErrNameEmpty = errors.New("Name is empty")
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024)

	var req dto.CreateUserRequest

	if err := decodeJSON(*r, &req); err != nil {
		writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}
	switch {
	case strings.TrimSpace(req.UserName) == "":
		writeError(w, "Invalid user name - empty", http.StatusBadRequest)

	case len(strings.TrimSpace(req.UserName)) > 30:
		writeError(w, "Name is too long", http.StatusBadRequest)

	case len(strings.TrimSpace(req.UserName)) < 3:
		writeError(w, "Name is too short", http.StatusBadRequest)
	}

	user := models.User{
		UserName: req.UserName,
	}

	createdUser, err := h.userService.Create(ctx, &user)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrUserAlreadyExists):
			writeError(w, "Name already exists", http.StatusConflict)
		default:
			log.Printf("create user %v", err)
			writeError(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	if err := writeJSON(w, http.StatusCreated, createdUser); err != nil {
		writeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
