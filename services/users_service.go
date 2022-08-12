package services

import (
	"github.com/devmaufh/go_books/domain/users"
	"github.com/devmaufh/go_books/utils/errors"
)

// CreateUser using the dao function
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser with the given id.
func GetUser(userId int64) (*users.User, *errors.RestError) {
	user := &users.User{}
	if err := user.Get(userId); err != nil {
		return nil, err
	}

	return user, nil
}
