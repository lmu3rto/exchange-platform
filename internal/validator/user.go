package validator

import (
	"errors"
	"strings"
	"unicode"
)

var (
	ErrNameLong    = errors.New("name is too long")
	ErrNameShort   = errors.New("name is too short")
	ErrNameEmpty   = errors.New("name is empty")
	ErrBodyEmpty   = errors.New("body is empty")
	ErrInvalidName = errors.New("invalid user name")
)

const (
	MaxLenName = 30
	MinLenName = 3
)

func UserName(name string) error {
	userName := strings.TrimSpace(name)

	switch {
	case userName == "":
		return ErrNameEmpty
	case len(userName) > MaxLenName:
		return ErrNameLong
	case len(userName) < MinLenName:
		return ErrNameShort
	}

	for _, symbol := range userName {
		if unicode.IsDigit(symbol) || symbol == '_' || !unicode.IsLetter(symbol) {
			return ErrInvalidName
		}
	}
	return nil
}
