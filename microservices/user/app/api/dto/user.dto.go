package dto

type CreateUserDto struct {
	Login     string `json:"login" validate:"required,min=3" example:"login"`
	Password  string `json:"password" validate:"required,min=8" example:"password"`
	FirstName string `json:"firstName" validate:"required" example:"firstName"`
	LastName  string `json:"lastName" validate:"required" example:"lastName"`
	Birthday  string `json:"birthday" validate:"required,date" example:"2000-01-01"`
}

type UpdateUserDto struct {
	Password  string `json:"password" validate:"omitempty,min=8" example:"password"`
	FirstName string `json:"firstName" example:"firstName"`
	LastName  string `json:"lastName" example:"lastName"`
	Birthday  string `json:"birthday" validate:"omitempty,date" example:"2000-01-01"`
}
