package users

import (
	"fmt"
	"github.com/devmaufh/go_books/database/postgres/users_db"
	"github.com/devmaufh/go_books/utils/date_utils"
	"github.com/devmaufh/go_books/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name,last_name,email,created_at,updated_at) VALUES(?,?,?,?,?)"
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
	user.CreatedAt = date_utils.GetNowString()
	user.UpdatedAt = date_utils.GetNowString()
	users_db.Client.Raw(queryInsertUser, user.FirstName, user.LastName, user.Email, user.CreatedAt, user.UpdatedAt).Scan(&user.Id)
	return nil
}
