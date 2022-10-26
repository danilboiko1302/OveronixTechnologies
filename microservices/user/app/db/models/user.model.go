package model

type User struct {
	Id        int    `json:"id" db:"id"`
	Login     string `json:"login" db:"login"`
	Password  string `json:"password" db:"password"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Birthday  string `json:"birthday" db:"birthday"`
}
