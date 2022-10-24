package http_controllers

import (
	"net/http"
	"user/app/api/dto"
	middlewares "user/app/controllers/middleware"

	"github.com/gin-gonic/gin"
)

func UserController(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/user", middlewares.Validator(dto.CreateUserDto{}), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"it works": "!!!"})
		})

		userGroup.PUT("/user/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"it works": "!!!"})
		})

		userGroup.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"it works": "!!!"})
		})

		userGroup.GET("/user/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"it works": "!!!"})
		})

		userGroup.DELETE("/user/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"it works": "!!!"})
		})
	}
}
