package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

// StartApplication Initialize the application.
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
