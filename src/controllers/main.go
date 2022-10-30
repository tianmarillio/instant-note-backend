package controllers

import (
	"github.com/tianmarillio/instant-note-backend/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = config.DB
}
