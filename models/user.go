package models

type User struct {
	Id       uint
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
