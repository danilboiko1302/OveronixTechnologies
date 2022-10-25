package services

import (
	"user/app/api/dto"
	model "user/app/db/models"
	"user/app/db/queries"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(data *dto.CreateUserDto) (*model.User, error) {
	hash, err := hashPassword(data.Password)

	if err != nil {
		return nil, err
	}

	data.Password = hash

	return queries.SQLSession.CreateUser(data)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
