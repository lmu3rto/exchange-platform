package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/lmu3rto/exchange-platform/internal/domain/errs"
	"github.com/lmu3rto/exchange-platform/internal/domain/models"
	"github.com/lmu3rto/exchange-platform/internal/handler/dto"
	"github.com/lmu3rto/exchange-platform/internal/validator"
)

const (
	MaxBytesMemory = 1024 * 1024
	FiveSeconds    = 5 * time.Second
)

// CreateUser godoc
// @Summary Create user
// @Description Creates a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "Create user request"
// @Success 201 {object} dto.CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
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
		if errors.Is(err, errs.ErrUserAlreadyExists) {
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

// GetUserByID godoc
// @Summary Get user by ID
// @Description Returns a user by ID
// @Tags users
// @Produce json
// @Param id path int64 true "User ID"
// @Success 200 {object} dto.GetByIDResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
		if errors.Is(err, errs.ErrUserNotFound) {
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

// GetUserByName godoc
// @Summary Get user by name
// @Description Returns a user by username
// @Tags users
// @Produce json
// @Param name query string true "Username"
// @Success 200 {object} dto.GetByNameResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/name [get]
func (h *UserHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, cancel := context.WithTimeout(ctx, FiveSeconds)
	defer cancel()

	req := dto.GetByNameRequest{
		UserName: strings.TrimSpace(r.URL.Query().Get("name")),
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

// UpdateUser godoc
// @Summary Update user
// @Description Updates user's username
// @Tags users
// @Accept json
// @Produce json
// @Param id path int64 true "User ID"
// @Param request body dto.UpdateRequest true "Update user request"
// @Success 200 {object} dto.UpdateResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [put]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
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
		if errors.Is(err, errs.ErrUserAlreadyExists) {
			writeError(w, "name already exists", http.StatusConflict)
			return
		}
		if errors.Is(err, errs.ErrUserNotFound) {
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

// DeleteUser godoc
// @Summary Delete user
// @Description Deletes a user by ID
// @Tags users
// @Produce json
// @Param id path int64 true "User ID"
// @Success 200 {object} dto.DeleteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
		if errors.Is(err, errs.ErrUserNotFound) {
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
