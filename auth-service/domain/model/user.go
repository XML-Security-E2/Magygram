package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"html"
	"log"
	"regexp"
)

type User struct {
	Id string `bson:"_id,omitempty"`
	Active bool `bson:"active"`
	Email string `bson:"email" validate:"required,email"`
	Password string `bson:"password"`
	Roles []Role `bson:"roles"`
}

type UserRequest struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type ActivatedRequest struct {
	Email string `json:"email"`
}

type PasswordChangeRequest struct {
	UserId string `json:"userId"`
	Password string `json:"password"`
	PasswordRepeat string `json:"passwordRepeat"`
}

func NewUser(userRequest *UserRequest) (*User, error) {
	hashAndSalt, err := HashAndSaltPasswordIfStrongAndMatching(userRequest.Password, userRequest.RepeatedPassword)
	if err != nil {
		return nil, err
	}
	return &User{Id: userRequest.Id,
		Active:   false,
		Email:    html.EscapeString(userRequest.Email),
		Password: hashAndSalt,
		Roles: []Role{{ Name: "user", Permissions: []Permission{{"create_post"}}}}}, err
}

func HashAndSaltPasswordIfStrongAndMatching(password string, repeatedPassword string) (string, error) {
	isMatching := password == repeatedPassword
	if !isMatching {
		return "", errors.New("passwords are not matching")
	}
	isWeak, _ := regexp.MatchString("^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[^!@#$%^&*(),.?\":{}|<>~'_+=]*)$", password)

	if isWeak {
		return "", errors.New("password must contain minimum eight characters, at least one capital letter, one number and one special character")
	}
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash), err
}
