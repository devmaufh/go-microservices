package app

import (
	"github.com/devmaufh/go_books/http/controllers/users"
	"github.com/devmaufh/go_books/http/middlewares"
)

// Map the Http routes.
func mapUrls() {
	router.GET("/users/:user_id", middlewares.RequestLoggerStart(), users.GetUser)
	//router.GET("/users/search", userss.SearchUser)
	router.POST("/users", middlewares.RequestLoggerStart(), users.CreateUser)
}
