package models

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Name    string
	Members []User `gorm:"many2many:chat_members;"`
}
