package http_controllers

import (
	"errors"
	"net/http"
	"sync"
	"time"
	"user/app/api/constants"
	"user/app/api/dto"
	middlewares "user/app/controllers/middleware"
	"user/app/services"
	voc "user/app/vocabulary"

	_ "user/docs"

	"github.com/gin-gonic/gin"
)

type Limit struct {
	mutex   *sync.Mutex
	current int
}

var limit *Limit = &Limit{
	mutex:   &sync.Mutex{},
	current: 0,
}

func (limit *Limit) add() error {
	limit.mutex.Lock()

	if limit.current == 0 {
		go func() {
			<-time.NewTimer(time.Second * 5).C
			limit.zero()
		}()
	}

	if limit.current == constants.LimitPerMinute {
		limit.mutex.Unlock()
		return errors.New(voc.LIMIT_PER_MINUTE_REACHED)
	}

	limit.current = limit.current + 1
	limit.mutex.Unlock()
	return nil
}

func (limit *Limit) zero() {
	limit.mutex.Lock()
	limit.current = 0
	limit.mutex.Unlock()
}

//!!!!!!!!
//controller was reworked, swagger needs one annotation per one func

// CreateUser             godoc
// @Summary      Create User
// @Description  Create User
// @Tags         user
// @Produce      json
// @Param request body dto.CreateUserDto true "data"
// @Success      200 {object} model.User
// @Failure      400  {string} string "Validation error"
// @Failure      500  {string} string  "Internal error"
// @Router       /user [post]
func CreateUser(router *gin.Engine) {
	router.POST("/users/user", middlewares.Validator(dto.CreateUserDto{}), func(c *gin.Context) {
		data := c.MustGet("validData").(*dto.CreateUserDto)

		response, err := services.CreateUser(data)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, response)
	})

}

// @Summary      Get User
// @Description  Get User
// @Tags         user
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200 {object} model.User
// @Failure      400  {string} string "not found"
// @Failure      500  {string} string  "Internal error"
// @Router       /user/{id} [get]
func GetUser(router *gin.Engine) {
	router.GET("/users/user/:id", func(c *gin.Context) {
		if err := limit.add(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		id := c.Param("id")

		response, err := services.GetUser(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, response)
	})
}

// @Summary      Get All Users
// @Description  Get All Users
// @Tags         user
// @Produce      json
// @Success      200 {array} model.User
// @Failure      500  {string} string  "Internal error"
// @Router       /user [get]
func GetUsers(router *gin.Engine) {
	router.GET("/users/user", func(c *gin.Context) {
		if err := limit.add(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		response, err := services.GetUsers()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"users": response})
	})
}

// @Summary      Delete User
// @Description  Delete User
// @Tags         user
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200 {object} model.User
// @Failure      400  {string} string "not found"
// @Failure      500  {string} string  "Internal error"
// @Router       /user/{id} [delete]
func DeleteUser(router *gin.Engine) {
	router.DELETE("/users/user/:id", func(c *gin.Context) {
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

// @Summary      Update User
// @Description  Update User
// @Tags         user
// @Produce      json
// @Param request body dto.UpdateUserDto true "data"
// @Success      200 {object} model.User
// @Failure      400  {string} string "Validation error"
// @Failure      500  {string} string  "Internal error"
// @Router       /user [put]
func UpdateUser(router *gin.Engine) {
	router.PUT("/users/user/:id", middlewares.Validator(dto.UpdateUserDto{}), func(c *gin.Context) {
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
}
