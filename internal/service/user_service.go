package service

import (
	"context"
	"errors"
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

func (s *UserService) Create(ctx context.Context, user *models.User) (*models.User, error) {

	name, err := s.repo.GetByName(ctx, strings.TrimSpace(user.UserName))

	if name != nil && !errors.Is(err, repository.ErrUserNotFound) {
		return nil, ErrUserAlreadyExists
	}

	return s.repo.Create(ctx, user)
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) GetByName(ctx context.Context, name string) (*models.User, error) {
	return s.repo.GetByName(ctx, name)

}

func (s *UserService) Update(ctx context.Context, user *models.User) (*models.User, error) {
	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, user *models.User) (*models.User, error) {
	return s.repo.Delete(ctx, user)
}
