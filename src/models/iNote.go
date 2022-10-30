package models

import "gorm.io/gorm"

type INote struct {
	gorm.Model
	ID      string
	Title   string
	Body    string
	FileUrl string
	Type    string // enum: "TEXT" || "FILE"

	UserID string
	// User   User
}
