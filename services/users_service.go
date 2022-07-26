package services

import "github.com/devmaufh/go_books/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
