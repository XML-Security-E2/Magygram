package model

type Permission struct {
	Id string `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"unique" json:"name"`
	Roles []Role `gorm:"many2many:role_permissions;" json:"-"`
}
