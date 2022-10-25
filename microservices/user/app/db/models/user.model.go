package model

type User struct {
	Id        int    `json:"id" db:"id"`
	Login     string `json:"login" db:"login"`
	Password  string `json:"password" db:"password"`
	FirstName string `json:"firstName" db:"firstName"`
	LastName  string `json:"lastName" db:"lastName"`
	Birthday  string `json:"birthday" db:"birthday"`
}
