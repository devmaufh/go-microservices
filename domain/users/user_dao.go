package users

import (
	"fmt"
	"github.com/devmaufh/go_books/utils/date_utils"
	"github.com/devmaufh/go_books/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get(userId int64) *errors.RestError {

	result := usersDB[userId]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found.", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedAt = result.CreatedAt
	user.UpdatedAt = result.UpdatedAt
	return nil
}

func (user *User) Save() *errors.RestError {
	if current := usersDB[user.Id]; current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("The email %s is already taken.", current.Email))
		}
		return errors.NewBadRequestError("The user already exists.")
	}

	user.CreatedAt = date_utils.GetNowString()
	user.UpdatedAt = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
