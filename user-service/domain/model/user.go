package model

import (
	"github.com/beevik/guid"
	"html"
)

type User struct {
	Id string `bson:"_id,omitempty"`
	Name  string `bson:"name" validate:"required,min=2"`
	Email string `bson:"email" validate:"required,email"`
	Surname string `bson:"surname" validate:"required,min=2"`
}

type UserRequest struct {
	Name  string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Password string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type ActivateLinkRequest struct {
	Email string `json:"email"`
}

type ChangeNewPasswordRequest struct {
	ResetPasswordId string `json:"resetPasswordId"`
	Password string `json:"password"`
	PasswordRepeat string `json:"passwordRepeat"`
}

func NewUser(userRequest *UserRequest) *User {
	return &User{Id: guid.New().String(),
		Name:     html.EscapeString(userRequest.Name),
		Surname:  html.EscapeString(userRequest.Surname),
		Email:    html.EscapeString(userRequest.Email)}
}