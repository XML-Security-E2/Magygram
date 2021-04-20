package model

type User struct {
	Id string `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
	Surname string `json:"surname"`
}
