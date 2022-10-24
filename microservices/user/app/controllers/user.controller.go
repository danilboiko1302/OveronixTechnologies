package http_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserController(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"it works": "!!!"})
		})
	}
}
