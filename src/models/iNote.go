package models

type INote struct {
	// gorm.Model

	ID      string `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	FileUrl string `json:"fileUrl"`
	Type    string `json:"type"` // enum: "TEXT" || "FILE"

	UserID string `json:"userId"`
	// User   User

	TimeModel
}



