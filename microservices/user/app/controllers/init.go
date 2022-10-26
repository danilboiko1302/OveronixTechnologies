package http_controllers

import (
	"github.com/gin-gonic/gin"
)

func InitControllers() *gin.Engine {
	router := gin.Default()

	CreateUser(router)
	UpdateUser(router)
	GetUser(router)
	GetUsers(router)
	DeleteUser(router)
	SwaggerController(router)
	return router
}
