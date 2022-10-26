package http_controllers

import (
	"net/http"
	"user/app/api/dto"
	middlewares "user/app/controllers/middleware"
	"user/app/services"

	"github.com/gin-gonic/gin"
)

func UserController(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/user", middlewares.Validator(dto.CreateUserDto{}), func(c *gin.Context) {
			data := c.MustGet("validData").(*dto.CreateUserDto)

			response, err := services.CreateUser(data)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			c.JSON(http.StatusOK, response)
		})

		userGroup.PUT("/user/:id", middlewares.Validator(dto.UpdateUserDto{}), func(c *gin.Context) {
			id := c.Param("id")
			data := c.MustGet("validData").(*dto.UpdateUserDto)
			response, err := services.UpdateUser(data, id)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			c.JSON(http.StatusOK, response)
		})

		userGroup.GET("/user", func(c *gin.Context) {
			response, err := services.GetUsers()

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			c.JSON(http.StatusOK, gin.H{"users": response})
		})

		userGroup.GET("/user/:id", func(c *gin.Context) {
			id := c.Param("id")

			response, err := services.GetUser(id)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			c.JSON(http.StatusOK, response)
		})

		userGroup.DELETE("/user/:id", func(c *gin.Context) {
			id := c.Param("id")

			response, err := services.DeleteUser(id)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			c.JSON(http.StatusOK, response)
		})
	}
}
