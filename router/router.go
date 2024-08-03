package router

import (
	"ip-data-api/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	router := gin.Default()
	router.GET("/", handlers.GetIpData)
	router.Run(":8080")
}
