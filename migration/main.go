package main

import (
	"github.com/tianmarillio/instant-note-backend/config"
	"github.com/tianmarillio/instant-note-backend/src/models"
)

func main() {
	config.LoadEnv()
	config.ConnectToDB()
	config.DB.AutoMigrate(
		&models.User{},
		&models.INote{},
	)
}
