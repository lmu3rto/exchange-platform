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

type CustomerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func scanCustomer(row pgx.Row) (*models.Customer, error) {
	var ex models.Customer

	err := row.Scan(
		&ex.ID,
		&ex.UserID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrCustomerNotFound
		}
		return nil, fmt.Errorf("failed scan customer repository: %w", err)
	}
	return &ex, nil
}

func (r *CustomerRepository) Create(ctx context.Context, id int64) (*models.Customer, error) {
	query := `
	INSERT INTO customers (user_id)
	VALUES($1)
	RETURNING
		id,
		user_id
	`

	return scanCustomer(r.db.QueryRow(ctx, query, id))
}

func (r *CustomerRepository) GetByID(ctx context.Context, id int64) (*models.Customer, error) {
	query := `
	SELECT id, user_id
	FROM customers
	WHERE user_id = $1
	`

	return scanCustomer(r.db.QueryRow(ctx, query, id))
}

func (r *CustomerRepository) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM customers
	WHERE user_id = $1
	RETURNING
		id,
		user_id
	`
	_, err := scanCustomer(r.db.QueryRow(ctx, query, id))
	return err
}
