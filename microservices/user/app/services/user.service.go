package services

import (
	"errors"
	"user/app/api/dto"
	model "user/app/db/models"
	"user/app/db/queries"
	voc "user/app/vocabulary"

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

func UpdateUser(data *dto.UpdateUserDto, id string) (*model.User, error) {
	newValues, err := checkUpdateValues(data)

	if err != nil {
		return nil, err
	}

	if len(newValues) == 0 {
		return nil, errors.New(voc.EMPTY_DATA_FOR_UPDATE_USER)
	}

	return queries.SQLSession.UpdateUser(newValues, id)
}

func checkUpdateValues(data *dto.UpdateUserDto) (map[string]string, error) {
	if data == nil {
		return nil, nil
	}
	var newValues map[string]string = make(map[string]string)

	if data.FirstName != "" {
		newValues["\"firstName\""] = data.FirstName
	}

	if data.LastName != "" {
		newValues["\"lastName\""] = data.LastName
	}

	if data.Birthday != "" {
		newValues["birthday"] = data.Birthday
	}

	if data.Password != "" {
		hash, err := hashPassword(data.Password)

		if err != nil {
			return nil, err
		}
		newValues["password"] = hash
	}

	return newValues, nil
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
