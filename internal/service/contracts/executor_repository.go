package contracts

import (
	"context"

	"github.com/lmu3rto/exchange-platform/internal/domain/models"
)

type ExecutorRepository interface {
	Create(ctx context.Context, id int64) (*models.Executor, error)
	GetByID(ctx context.Context, id int64) (*models.Executor, error)
	Delete(ctx context.Context, id int64) error
}
