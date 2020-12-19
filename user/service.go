package user

import (
	"errors"
	"strings"
)

type Service interface {
	AddUser(string) (string, error)
}

type User struct{}

func (User) AddUser(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for Service.
type ServiceMiddleware func(Service) Service
