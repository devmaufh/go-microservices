package users

import (
	"github.com/devmaufh/go_books/domain/users"
	"github.com/devmaufh/go_books/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "todo")
}
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: handle the error.
		return
	}
	createdUser, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle the error.
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "todo")
}
