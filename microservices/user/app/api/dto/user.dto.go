package dto

type CreateUserDto struct {
	Login     string `json:"login" validate:"required,min=3" example:"login"`
	Password  string `json:"password" validate:"required,min=8" example:"password"`
	FirstName string `json:"first_name" validate:"required" example:"first_name"`
	LastName  string `json:"last_name" validate:"required" example:"last_name"`
	Birthday  string `json:"birthday" validate:"required,date" example:"2000-01-01"`
}

type UpdateUserDto struct {
	Password  string `json:"password" validate:"omitempty,min=8" example:"password"`
	FirstName string `json:"first_name" example:"first_name"`
	LastName  string `json:"last_name" example:"last_name"`
	Birthday  string `json:"birthday" validate:"omitempty,date" example:"2000-01-01"`
}
