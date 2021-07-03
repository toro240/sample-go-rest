package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp - Start our application
func StartApp() {
	mapUrls()

	log.Println("Start App...")
	router.Run(":8080")
}
