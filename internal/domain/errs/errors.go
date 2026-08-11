package errs

import (
	"errors"
)

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrExecutorNotFound      = errors.New("executor not found")
	ErrCustomerNotFound      = errors.New("customer not found")
	ErrCustomerAlreadyExists = errors.New("customer already exists")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrExecutorAlreadyExists = errors.New("executor already exists")
)
