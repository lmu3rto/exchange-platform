package service

import (
	"context"
	"errors"

	"github.com/lmu3rto/exchange-platform/internal/domain/errs"
	"github.com/lmu3rto/exchange-platform/internal/domain/models"
	"github.com/lmu3rto/exchange-platform/internal/service/contracts"
)

type ExecutorService struct {
	exrepo contracts.ExecutorRepository
	usrepo contracts.UserRepository
}

func NewExecutorService(es contracts.ExecutorRepository, us contracts.UserRepository) *ExecutorService {
	return &ExecutorService{
		exrepo: es,
		usrepo: us,
	}
}

func (s *ExecutorService) Create(ctx context.Context, id int64) (*models.Executor, error) {
	_, err := s.usrepo.GetByID(ctx, id)

	if errors.Is(err, errs.ErrUserNotFound) {
		return nil, errs.ErrUserNotFound
	}

	ex, err := s.exrepo.GetByID(ctx, id)

	if ex != nil {
		return nil, errs.ErrExecutorAlreadyExists
	}

	return s.exrepo.Create(ctx, id)
}

func (s *ExecutorService) GetByID(ctx context.Context, id int64) (*models.Executor, error) {
	ex, err := s.exrepo.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, errs.ErrCustomerNotFound) {
			return nil, errs.ErrCustomerNotFound
		}
		return nil, err
	}

	return ex, nil	
}

func (s *ExecutorService) Delete(ctx context.Context, id int64) error {
	err := s.exrepo.Delete(ctx, id)

	if err != nil {
		if errors.Is(err, errs.ErrExecutorNotFound) {
			return errs.ErrExecutorNotFound
		}
	}
	return nil
}
