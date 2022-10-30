package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string
	Email    string `gorm:"uniqueIndex"`
	Password string
}
