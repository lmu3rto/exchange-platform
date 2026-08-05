package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/lmu3rto/exchange-platform/internal/domain/models"
	"github.com/lmu3rto/exchange-platform/internal/handler/dto"
	"github.com/lmu3rto/exchange-platform/internal/service"
	"github.com/lmu3rto/exchange-platform/internal/validator"
)

const (
	MaxBytesMemory = 1024 * 1024
	FiveSeconds    = 5 * time.Second
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, cancel := context.WithTimeout(ctx, FiveSeconds)
	defer cancel()

	r.Body = http.MaxBytesReader(w, r.Body, MaxBytesMemory)

	var req dto.CreateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user := models.User{
		UserName: req.UserName,
	}

	if err := validator.UserName(user.UserName); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := h.userService.Create(ctx, &user)

	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			writeError(w, "Name already exists", http.StatusConflict)
			return
		}
		h.logger.Error(
			"create user failed",
			"error", err,
		)
		writeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	res := dto.CreateUserResponse{
		UserName:  createdUser.UserName,
		CreatedAt: createdUser.CreatedAt,
	}

	if err := writeJSON(w, http.StatusCreated, res); err != nil {
		h.logger.Error(
			"write json response failed",
			"error", err,
		)
		return
	}
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, cancel := context.WithTimeout(ctx, FiveSeconds)
	defer cancel()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeError(w, "invalid user id", http.StatusBadRequest)
		return
	}

	getUser, err := h.userService.GetByID(ctx, int64(id))

	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			writeError(w, "user not found", http.StatusNotFound)
		}
		h.logger.Error(
			"get user by id failed",
			"error", err,
		)
		writeError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := dto.GetByIDResponse{
		UserName:  getUser.UserName,
		CreatedAt: getUser.CreatedAt,
	}

	if err := writeJSON(w, http.StatusOK, res); err != nil {
		h.logger.Error(
			"write json response failed",
			"error", err,
		)
		return
	}

}

func (h *Handler) GetByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, cancel := context.WithTimeout(ctx, FiveSeconds)
	defer cancel()

	r.Body = http.MaxBytesReader(w, r.Body, MaxBytesMemory)

	var req dto.GetByNameRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := validator.UserName(req.UserName); err != nil {
		writeError(w, "invalid user name", http.StatusBadRequest)
		return
	}

	getUser, err := h.userService.GetByName(ctx, req.UserName)

	if err != nil {
		h.logger.Error(
			"get user by name failed",
			"error", err,
		)
		writeError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := dto.GetByNameResponse{
		UserName: getUser.UserName,
	}

	if err := writeJSON(w, http.StatusOK, res); err != nil {
		h.logger.Error(
			"write json response failed",
			"error", err,
		)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, cancel := context.WithTimeout(ctx, FiveSeconds)
	defer cancel()

	r.Body = http.MaxBytesReader(w, r.Body, MaxBytesMemory)

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeError(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var req dto.UpdateRequest

	if err := decodeJSON(r, &req); err != nil {
		writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := validator.UserName(req.UserName); err != nil {
		writeError(w, "invalid user name", http.StatusBadRequest)
		return
	}

	user := models.User{
		ID:       int64(id),
		UserName: req.UserName,
	}

	updateName, err := h.userService.Update(ctx, &user)

	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			writeError(w, "name already exists", http.StatusConflict)
			return
		}
		if errors.Is(err, service.ErrUserNotFound) {
			writeError(w, "user not found by id", http.StatusNotFound)
			return
		}
		h.logger.Error(
			"update user failed",
			"error", err,
		)
		writeError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := dto.UpdateResponse{
		UserName: updateName.UserName,
	}

	if err := writeJSON(w, http.StatusOK, res); err != nil {
		h.logger.Error(
			"write json response failed",
			"error", err,
		)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, cancel := context.WithTimeout(ctx, FiveSeconds)
	defer cancel()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeError(w, "invalid user id", http.StatusBadRequest)
		return
	}

	deletedUser, err := h.userService.Delete(ctx, int64(id))

	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			writeError(w, "user not found by id", http.StatusNotFound)
			return
		}
		h.logger.Error(
			"delete user failed",
			"error", err,
		)
		writeError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, http.StatusOK, deletedUser); err != nil {
		h.logger.Error(
			"write json response failed",
			"error", err,
		)
		return
	}

}
