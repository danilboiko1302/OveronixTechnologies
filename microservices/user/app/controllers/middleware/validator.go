package middlewares

import (
	"net/http"
	"user/app/api/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)

func Validator(i interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data interface{}
		switch i.(type) {
		case dto.CreateUserDto:
			data = &dto.CreateUserDto{}

		default:
			context.JSON(http.StatusInternalServerError, gin.H{"error": "dto type is invalid"})
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if err := context.Bind(data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		v := validator.New()

		if err := v.Struct(data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		conform := modifiers.New()

		if err := conform.Struct(context, data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		context.Set("validData", data)
		context.Next()
	}
}
