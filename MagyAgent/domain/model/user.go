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
	Roles []Role `gorm:"many2many:user_roles;"`
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
				 Password: HashAndSaltPassword(userRequest.Password),
			     Roles: []Role{{ Id: "7a753a24-5a20-4021-a3e0-0afdf3744675", Name: "user"}}}
}

func HashAndSaltPassword(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
