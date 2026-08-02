package handler

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/lmu3rto/exchange-platform/internal/domain/models"
	"github.com/lmu3rto/exchange-platform/internal/handler/dto"
	"github.com/lmu3rto/exchange-platform/internal/service"
)

var (
	ErrNameLong  = errors.New("Name is too long")
	ErrNameShort = errors.New("Name is too short")
	ErrNameEmpty = errors.New("Name is empty")
	ErrBodyEmpty = errors.New("body is empty")
)

const (
	MaxBytesMemory = 1024 * 1024
	MaxLenName     = 30
	MinLenName     = 3
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.Body = http.MaxBytesReader(w, r.Body, MaxBytesMemory)

	var req dto.CreateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	userName := strings.TrimSpace(req.UserName)
	switch {
	case userName == "":
		writeError(w, ErrNameEmpty.Error(), http.StatusBadRequest)
		return
	case len(userName) > MaxLenName:
		writeError(w, ErrNameLong.Error(), http.StatusBadRequest)
		return
	case len(userName) < MinLenName:
		writeError(w, ErrNameShort.Error(), http.StatusBadRequest)
		return
	}

	user := models.User{
		UserName: req.UserName,
	}

	createdUser, err := h.userService.Create(ctx, &user)
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			writeError(w, "Name already exists", http.StatusConflict)
			return
		}
		log.Printf("create user error: %v", err)
		writeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, http.StatusCreated, createdUser); err != nil {
		log.Printf("failed to write json response: %v", err)
	}
}
