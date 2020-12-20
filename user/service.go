package user

import (
	"errors"
	"math/rand"
)

type Service interface {
	AddUser(string, string) (int, error)
}

type User struct{}

func (User) AddUser(firstName, lastName string) (int, error) {
	if firstName == "" || lastName == "" {
		return -1, ErrEmpty
	}
	s := rand.NewSource(100)
	r := rand.New(s)
	return r.Int(), nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty inputs")

// ServiceMiddleware is a chainable behavior modifier for Service.
type ServiceMiddleware func(Service) Service
