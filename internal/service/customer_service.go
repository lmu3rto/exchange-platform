package service

import (
	"context"
	"errors"

	"github.com/lmu3rto/exchange-platform/internal/domain/errs"
	"github.com/lmu3rto/exchange-platform/internal/domain/models"
	"github.com/lmu3rto/exchange-platform/internal/service/contracts"
)
type CustomerService struct {
	csrepo contracts.CustomerRepository
	usrepo contracts.UserRepository
}

func NewCustomerService(es contracts.CustomerRepository, us contracts.UserRepository) *CustomerService {
	return &CustomerService{
		csrepo: es,
		usrepo: us,
	}
}

func (s *CustomerService) Create(ctx context.Context, id int64) (*models.Customer, error) {
	_, err := s.usrepo.GetByID(ctx, id)

	if errors.Is(err, errs.ErrUserNotFound) {
		return nil, errs.ErrUserNotFound
	}

	cs, err := s.csrepo.GetByID(ctx, id)

	if cs != nil {
		return nil, errs.ErrCustomerAlreadyExists
	}

	return s.csrepo.Create(ctx, id)
}

func (s *CustomerService) GetByID(ctx context.Context, id int64) (*models.Customer, error) {
	cs, err := s.csrepo.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, errs.ErrCustomerNotFound) {
			return nil, errs.ErrCustomerNotFound
		}
		return nil, err
	}

	return cs, nil
}

func (s *CustomerService) Delete(ctx context.Context, id int64) (error) {
	err := s.csrepo.Delete(ctx, id)

	if err != nil {
		if errors.Is(err, errs.ErrCustomerNotFound) {
			return errs.ErrCustomerNotFound
		}
	}
	return nil
}