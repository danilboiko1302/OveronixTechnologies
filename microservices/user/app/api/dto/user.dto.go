package dto

import "time"

type CreateUserDto struct {
	Login     string `json:"login" validate:"required,min=3"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Birthday  string `json:"birthday" validate:"required,date"`
}

type UpdateUserDto struct {
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Birthday  time.Time `json:"birthday"`
}
