package services

import (
	"user/app/api/dto"
	model "user/app/db/models"
	"user/app/db/queries"

	"golang.org/x/crypto/bcrypt"
)

func GetUsers() ([]model.User, error) {
	return queries.SQLSession.GetUsers()
}

func DeleteUser(id string) (*model.User, error) {
	return queries.SQLSession.DeleteUser(id)
}

func GetUser(id string) (*model.User, error) {
	return queries.SQLSession.GetUser(id)
}

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
