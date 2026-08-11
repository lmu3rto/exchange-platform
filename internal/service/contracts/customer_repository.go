package contracts

import (
	"context"

	"github.com/lmu3rto/exchange-platform/internal/domain/models"
)

type CustomerRepository interface {
	Create(ctx context.Context, id int64) (*models.Customer, error)
	GetByID(ctx context.Context, id int64) (*models.Customer, error)
	Delete(ctx context.Context, id int64) error
}
