package app

import "github.com/devmaufh/go_books/controllers/users"

func mapUrls() {
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", userss.SearchUser)
	router.POST("/users", users.CreateUser)
}
