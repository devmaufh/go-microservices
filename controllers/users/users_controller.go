package users

import (
	"github.com/devmaufh/go_books/domain/users"
	"github.com/devmaufh/go_books/services"
	"github.com/devmaufh/go_books/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid json")
		c.JSON(restError.StatusCode, restError)
		return
	}
	createdUser, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restError := errors.NewBadRequestError("The user_id must be a valid integer.")
		c.JSON(restError.StatusCode, restError)
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
