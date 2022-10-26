package middlewares

import (
	"fmt"
	"net/http"
	"regexp"
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
		case dto.UpdateUserDto:
			data = &dto.UpdateUserDto{}

		default:
			context.JSON(http.StatusInternalServerError, gin.H{"error": "dto type is invalid"})
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := context.Bind(data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		v := validator.New()

		_ = v.RegisterValidation("date", func(fl validator.FieldLevel) bool {
			// match, err := regexp.MatchString("(?:0[1-9]|[12][0-9]|3[01])[-](?:0[1-9]|1[012])[-](?:19[0-9]{2}|20[01][0-9]|2020)", fl.Field().String())
			match, err := regexp.MatchString("(?:19[0-9]{2}|20[01][0-9]|2020)[-](?:0[1-9]|1[012])[-](?:0[1-9]|[12][0-9]|3[01])", fl.Field().String())

			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				context.AbortWithStatus(http.StatusInternalServerError)
				return true
			}

			return match
		})

		if err := v.Struct(data); err != nil {

			if err.Error() == "Key: 'CreateUserDto.Birthday' Error:Field validation for 'Birthday' failed on the 'date' tag" {
				context.JSON(http.StatusBadRequest, gin.H{"error": "Key: 'CreateUserDto.Birthday' Error:Field validation for 'Birthday' failed on the 'date' tag\nPlease use YYYY-MM-DD"})
				context.AbortWithStatus(http.StatusBadRequest)
				return
			}

			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		conform := modifiers.New()

		if err := conform.Struct(context, data); err != nil {
			fmt.Println("2")
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		context.Set("validData", data)
		context.Next()
	}
}
