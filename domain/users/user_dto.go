package users

import (
	"github.com/devmaufh/go_books/utils/errors"
	"net/mail"
	"strings"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return errors.NewBadRequestError("Invalid user email.")
	}
	return nil
}
