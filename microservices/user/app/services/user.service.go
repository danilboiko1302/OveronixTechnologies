package services

import (
	"user/app/api/dto"
	model "user/app/db/models"
	"user/app/db/queries"
)

func CreateUser(data *dto.CreateUserDto) (*model.User, error) {
	return queries.SQLSession.CreateUser(data)
}
