package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/lmu3rto/exchange-platform/internal/domain/models"
	"github.com/lmu3rto/exchange-platform/internal/repository"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

var (
	ErrUserAlreadyExists = errors.New("User already exists")
)

func wrap(s string, err error) error {
	if err != nil {
		return fmt.Errorf("user service %s: - %w", s, err)
	}
	return nil
}

func (s *UserService) Create(ctx context.Context, user *models.User) (*models.User, error) {

	name, err := s.repo.GetByName(ctx, strings.TrimSpace(user.UserName))

	if name != nil && !errors.Is(err, repository.ErrUserNotFound) {
		return nil, ErrUserAlreadyExists
	}

	res, err := s.repo.Create(ctx, user)

	return res, wrap("create", err)
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*models.User, error) {
	res, err := s.repo.GetByID(ctx, id)

	return res, wrap("get by id", err)
}

func (s *UserService) GetByName(ctx context.Context, name string) (*models.User, error) {
	res, err := s.repo.GetByName(ctx, name)

	return res, wrap("get by name", err)

}

func (s *UserService) Update(ctx context.Context, user *models.User) (*models.User, error) {
	res, err := s.repo.Update(ctx, user)

	return res, wrap("update", err)
}

func (s *UserService) Delete(ctx context.Context, user *models.User) (*models.User, error) {
	res, err := s.repo.Delete(ctx, user)

	return res, wrap("delete", err)
}
