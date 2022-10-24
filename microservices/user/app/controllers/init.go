package http_controllers

import (
	"github.com/gin-gonic/gin"
)

func InitControllers() *gin.Engine {
	router := gin.Default()

	UserController(router)

	return router
}
