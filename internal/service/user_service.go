package service

import (
	"context"
	"github.com/lmu3rto/exchange-platform/internal/repository"
	"github.com/lmu3rto/exchange-platform/internal/domain/models"

)
type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService {
		repo: r,
	}
}

func (s *UserService) Create(ctx context.Context, user models.User) (*models.User, error) {
	return s.repo.Create(ctx, &user)
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






