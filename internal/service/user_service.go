package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/lmu3rto/exchange-platform/internal/domain/errs"
	"github.com/lmu3rto/exchange-platform/internal/domain/models"
	"github.com/lmu3rto/exchange-platform/internal/service/contracts"
)

type UserService struct {
	repo contracts.UserRepository
}

func NewUserService(r contracts.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func wrap(s string, err error) error {
	if err != nil {
		return fmt.Errorf("user service %s: - %w", s, err)
	}
	return nil
}

func (s *UserService) Create(ctx context.Context, user *models.User) (*models.User, error) {

	name, err := s.repo.GetByName(ctx, user.UserName)

	if name != nil && !errors.Is(err, errs.ErrUserNotFound) {
		return nil, errs.ErrUserAlreadyExists
	}

	res, err := s.repo.Create(ctx, user)

	return res, wrap("create", err)
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*models.User, error) {

	res, err := s.repo.GetByID(ctx, id)

	if res == nil {
		return nil, err
	}

	if res.ID == 0 && errors.Is(err, errs.ErrUserNotFound) {
		return nil, errs.ErrUserNotFound
	}

	return res, wrap("get by id", err)
}

func (s *UserService) GetByName(ctx context.Context, name string) (*models.User, error) {
	res, err := s.repo.GetByName(ctx, name)

	return res, wrap("get by name", err)

}

func (s *UserService) Update(ctx context.Context, user *models.User) (*models.User, error) {

	ex, err := s.repo.GetByName(ctx, user.UserName)

	if ex != nil && !errors.Is(err, errs.ErrUserNotFound) {
		return nil, errs.ErrUserAlreadyExists
	}

	res, err := s.repo.Update(ctx, user)

	return res, wrap("update", err)
}

func (s *UserService) Delete(ctx context.Context, id int64) (*models.User, error) {
	res, err := s.repo.Delete(ctx, id)

	return res, wrap("delete", err)
}
