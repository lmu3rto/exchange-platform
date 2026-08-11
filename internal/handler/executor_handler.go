package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/lmu3rto/exchange-platform/internal/domain/errs"
)

// CreateExecutor godoc
// @Summary Create executor
// @Description Creates an executor role for the specified user
// @Tags executors
// @Accept json
// @Produce json
// @Param id path int64 true "User ID"
// @Success 201 {object} models.Executor
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 409 {object} ErrorResponse "Executor already exists"
// @Failure 500 {object} ErrorResponse
// @Router /users/{id}/executor [post]
func (h *ExecutorHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, cancel := context.WithTimeout(ctx, FiveSeconds)

	defer cancel()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeError(w, "invalid user id", http.StatusBadRequest)
		return
	}

	ex, err := h.executorService.Create(ctx, int64(id))

	if err != nil {
		if errors.Is(err, errs.ErrExecutorAlreadyExists) {
			writeError(w, "executor already exists", http.StatusConflict)
			return
		}
		if errors.Is(err, errs.ErrUserNotFound) {
			writeError(w, "user by id not found", http.StatusConflict)
			return
		}
		h.logger.Error("get executor by id error", "error", err)
		writeError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, http.StatusCreated, ex); err != nil {
		h.logger.Error("failed write json response", "error", err)
		return
	}
}

// GetExecutorByID godoc
// @Summary Get executor
// @Description Returns executor by user ID
// @Tags executors
// @Accept json
// @Produce json
// @Param id path int64 true "User ID"
// @Success 200 {object} models.Executor
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse "Executor not found"
// @Failure 500 {object} ErrorResponse
// @Router /users/{id}/executor [get]
func (h *ExecutorHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, cancel := context.WithTimeout(ctx, FiveSeconds)
	defer cancel()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeError(w, "invalid user id", http.StatusBadRequest)
		return
	}

	ex, err := h.executorService.GetByID(ctx, int64(id))

	if err != nil {
		if errors.Is(err, errs.ErrExecutorNotFound) {
			writeError(w, "executor by id not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, errs.ErrExecutorAlreadyExists) {
			writeError(w, "user already executor", http.StatusConflict)
			return
		}
		h.logger.Error("failed get executor by id", "error", err)
		writeError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, http.StatusOK, ex); err != nil {
		h.logger.Error("failed to write json response", "error", err)
		return
	}
}

// DeleteExecutor godoc
// @Summary Delete executor
// @Description Deletes executor role from the specified user
// @Tags executors
// @Accept json
// @Produce json
// @Param id path int64 true "User ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse "Executor not found"
// @Failure 500 {object} ErrorResponse
// @Router /users/{id}/executor [delete]
func (h *ExecutorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, cancel := context.WithTimeout(ctx, FiveSeconds)
	defer cancel()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeError(w, "invalid executor id", http.StatusBadRequest)
		return
	}

	if err := h.executorService.Delete(ctx, int64(id)); err != nil {
		if errors.Is(err, errs.ErrExecutorNotFound) {
			writeError(w, "executor not found", http.StatusNotFound)
			return
		}
		writeError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, http.StatusNoContent, nil); err != nil {
		h.logger.Error("failed write json response", "error", err)
		return
	}
}
