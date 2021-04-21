package model

import (
	"github.com/beevik/guid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Id string `gorm:"primaryKey"`
	Active bool
	Name  string
	Email string `gorm:"unique"`
	Password string
	Surname string
	Role string
}

type UserRequest struct {
	Name  string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func NewUser(userRequest *UserRequest) *User {
	return &User{Id: guid.New().String(),
				 Active: false,
				 Name: userRequest.Name,
				 Surname: userRequest.Surname,
				 Email: userRequest.Email,
				 Password: hashAndSaltPassword(userRequest.Password),
			     Role: "admin"}
}

func hashAndSaltPassword(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
