package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/lmu3rto/exchange-platform/internal/domain/errs"
)

// CreateCustomer godoc
// @Summary Create customer
// @Description Creates an customer role for the specified user
// @Tags customer
// @Accept json
// @Produce json
// @Param id path int64 true "User ID"
// @Success 201 {object} models.Customer
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 409 {object} ErrorResponse "Customer already exists"
// @Failure 500 {object} ErrorResponse
// @Router /users/{id}/customer [post]
func (h *CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, cancel := context.WithTimeout(ctx, FiveSeconds)

	defer cancel()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeError(w, "invalid user id", http.StatusBadRequest)
		return
	}

	ex, err := h.customerService.Create(ctx, int64(id))

	if err != nil {
		if errors.Is(err, errs.ErrCustomerAlreadyExists) {
			writeError(w, "user already customer", http.StatusConflict)
			return
		}
		if errors.Is(err, errs.ErrUserNotFound) {
			writeError(w, "user by id not found", http.StatusConflict)
			return
		}
		h.logger.Error("failed create customer by id", "error", err)
		writeError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, http.StatusCreated, ex); err != nil {
		h.logger.Error("failed write json response", "error", err)
		return
	}
}

// GetCustomerByID godoc
// @Summary Get customer
// @Description Returns customer by user ID
// @Tags customer
// @Accept json
// @Produce json
// @Param id path int64 true "User ID"
// @Success 200 {object} models.Customer
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse "Customer not found"
// @Failure 500 {object} ErrorResponse
// @Router /users/{id}/customer [get]
func (h *CustomerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, cancel := context.WithTimeout(ctx, FiveSeconds)
	defer cancel()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeError(w, "invalid user id", http.StatusBadRequest)
		return
	}

	cs, err := h.customerService.GetByID(ctx, int64(id))

	if err != nil {
		if errors.Is(err, errs.ErrCustomerNotFound) {
			writeError(w, "customer by id not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, errs.ErrCustomerAlreadyExists) {
			writeError(w, "user already customer", http.StatusConflict)
			return
		}
		h.logger.Error("failed get customer by id", "error", err)
		writeError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, http.StatusOK, cs); err != nil {
		h.logger.Error("failed to write json response", "error", err)
		return
	}

}

// DeleteCustomer godoc
// @Summary Delete customer
// @Description Deletes customer role from the specified user
// @Tags customer
// @Accept json
// @Produce json
// @Param id path int64 true "User ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse "Customer not found"
// @Failure 500 {object} ErrorResponse
// @Router /users/{id}/customer [delete]
func (h *CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, cancel := context.WithTimeout(ctx, FiveSeconds)
	defer cancel()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeError(w, "invalid customer id", http.StatusBadRequest)
		return
	}

	if err := h.customerService.Delete(ctx, int64(id)); err != nil {
		if errors.Is(err, errs.ErrCustomerNotFound) {
			writeError(w, "customer not found", http.StatusNotFound)
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
