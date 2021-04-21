package model

import (
	"github.com/beevik/guid"
)

type User struct {
	Id string `gorm:"primaryKey"`
	Active bool
	Name  string
	Email string `gorm:"unique"`
	Password string
	Surname string
}

type UserRequest struct {
	Name  string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
}

func NewUser(userRequest *UserRequest) *User {
	return &User{Id: guid.New().String(), Active: false, Name: userRequest.Name, Surname: userRequest.Surname, Email: userRequest.Email, Password: userRequest.Password}
}
