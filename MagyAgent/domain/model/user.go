package model

import (
	"errors"
	"fmt"
	"github.com/beevik/guid"
	_ "github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

type User struct {
	Id string `gorm:"primaryKey"`
	Active bool `gorm:"required"`
	Name  string
	Email string `gorm:"unique" validate:"required,email"`
	Password string
	Surname string
	Roles []Role `gorm:"many2many:user_roles;"`
}

type UserRequest struct {
	Name  string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Password string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

func NewUser(userRequest *UserRequest) (*User,error) {
	hashAndSalt, err := HashAndSaltPasswordIfStrong(userRequest.Password)
	if err != nil {
		return nil, err
	}
	return &User{Id: guid.New().String(),
				 Active: false,
				 Name: userRequest.Name,
				 Surname: userRequest.Surname,
				 Email: userRequest.Email,
				 Password: hashAndSalt,
			     Roles: []Role{{ Id: "7a753a24-5a20-4021-a3e0-0afdf3744675", Name: "user"}}}, err
}

func HashAndSaltPasswordIfStrong(password string) (string, error) {

	isWeak, _ := regexp.MatchString("^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*)$", password)
	fmt.Println(isWeak)
	if isWeak {
		return password, errors.New("password must contain minimum eight characters, at least one capital letter and one number")
	}
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash), err
}
