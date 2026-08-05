package dto

import (
	"time"
)

type CreateUserRequest struct {
	UserName string `json:"user_name"`
}

type CreateUserResponse struct {
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}

type GetByIDResponse struct {
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type GetByNameResponse struct {
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type GetByNameRequest struct {
	UserName string `json:"user_name"`
}

type UpdateRequest struct {
	UserName string `json:"user_name"`
}

type UpdateResponse struct {
	UserName string `json:"user_name"`
}

type DeleteResponse struct {
	UserName  string    `json:"user_name"`
	DeletedAt time.Time `json:"deleted_at"`
}
