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

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func userScan(row pgx.Row) (*models.User, error) {
	var user models.User

	err := row.Scan(
		&user.ID,
		&user.UserName,
		&user.Balance,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed scan user repository: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	query := `
		INSERT INTO users (user_name)
		VALUES($1)
		RETURNING 
		  id,
    	user_name,
    	balance,
    	created_at,
    	updated_at,
    	deleted_at
		`
	return userScan(r.db.QueryRow(ctx, query, user.UserName))

}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := `
	SELECT 
	    id,
    	user_name,
    	balance,
    	created_at,
    	updated_at,
    	deleted_at
	FROM users
	WHERE id = ($1)
	`

	return userScan(r.db.QueryRow(ctx, query, id))
}

func (r *UserRepository) GetByName(ctx context.Context, name string) (*models.User, error) {
	query := `
		SELECT
    	id,
    	user_name,
    	balance,
    	created_at,
    	updated_at,
    	deleted_at
		FROM users
		WHERE user_name = $1
	`

	return userScan(r.db.QueryRow(ctx, query, name))
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) (*models.User, error) {
	query := `
	UPDATE users
	SET 
		user_name = ($1),
		balance = ($2)

	WHERE id = ($3)
	RETURNING 
	    id,
    	user_name,
    	balance,
    	created_at,
    	updated_at,
    	deleted_at
	`
	return userScan(r.db.QueryRow(ctx, query, user.UserName, user.Balance, user.ID))
}

func (r *UserRepository) Delete(ctx context.Context, id int64) (*models.User, error) {
	query := `
		UPDATE users
		SET deleted_at = NOW()
		WHERE id = $1
		RETURNING
    id,
    user_name,
    balance,
    created_at,
    updated_at,
    deleted_at
	`
	return userScan(r.db.QueryRow(ctx, query, id))
}
