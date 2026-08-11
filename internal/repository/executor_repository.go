package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmu3rto/exchange-platform/internal/domain/errs"
	"github.com/lmu3rto/exchange-platform/internal/domain/models"
)

type ExecutorRepository struct {
	db *pgxpool.Pool
}

func NewExecutorRepository(db *pgxpool.Pool) *ExecutorRepository {
	return &ExecutorRepository{
		db: db,
	}
}

func scanExecutor(row pgx.Row) (*models.Executor, error) {
	var ex models.Executor

	err := row.Scan(
		&ex.ID,
		&ex.UserID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrExecutorNotFound
		}
		return nil, fmt.Errorf("failed scan executor repository: %w", err)
	}
	return &ex, nil
}

func (r *ExecutorRepository) Create(ctx context.Context, id int64) (*models.Executor, error) {
	query := `
	INSERT INTO executors (user_id)
	VALUES($1)
	RETURNING
		id,
		user_id
	`

	return scanExecutor(r.db.QueryRow(ctx, query, id))
}

func (r *ExecutorRepository) GetByID(ctx context.Context, id int64) (*models.Executor, error) {
	query := `
	SELECT id, user_id
	FROM executors
	WHERE user_id = $1
	`

	return scanExecutor(r.db.QueryRow(ctx, query, id))
}

func (r *ExecutorRepository) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM executors
	WHERE user_id = $1
	RETURNING
		id,
		user_id
	`
	_, err := scanExecutor(r.db.QueryRow(ctx, query, id))
	return err
}
