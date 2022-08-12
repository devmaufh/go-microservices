package users

import (
	"github.com/devmaufh/go_books/database/postgres/users_db"
	"github.com/devmaufh/go_books/utils/date_utils"
	"github.com/devmaufh/go_books/utils/errors"
)

const (
	queryInsertUser     = "INSERT INTO users(first_name,last_name,email,created_at,updated_at) VALUES(?,?,?,?,?) RETURNING ID"
	querySelectUserById = "SELECT * FROM users where id = ?"
)

//Get the user with the given id.
func (user *User) Get(userId int64) *errors.RestError {
	result := users_db.Client.Raw(querySelectUserById, userId).Scan(&user)

	if result.Error != nil {
		return errors.NewInternalServerError("Internal server error.")
	}
	return nil
}

//Save the user in the users_db.
func (user *User) Save() *errors.RestError {
	user.CreatedAt = date_utils.GetNowString()
	user.UpdatedAt = date_utils.GetNowString()
	users_db.Client.Raw(queryInsertUser, user.FirstName, user.LastName, user.Email, user.CreatedAt, user.UpdatedAt).Scan(&user.Id)
	return nil
}
