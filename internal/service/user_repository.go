package service

import (
	"context"

	"github.com/lmu3rto/exchange-platform/internal/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByName(ctx context.Context, name string) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, user *models.User) (*models.User, error)
}
