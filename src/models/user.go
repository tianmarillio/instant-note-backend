package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`

	TimeModel
}
