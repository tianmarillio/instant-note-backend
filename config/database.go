package config

import (
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var sb strings.Builder
	var err error

	sb.WriteString("host=")
	sb.WriteString(os.Getenv("DB_HOST"))
	sb.WriteString(" ")
	sb.WriteString("user=")
	sb.WriteString(os.Getenv("DB_USERNAME"))
	sb.WriteString(" ")
	sb.WriteString("password=")
	sb.WriteString(os.Getenv("DB_PASSWORD"))
	sb.WriteString(" ")
	sb.WriteString("dbname=")
	sb.WriteString(os.Getenv("DB_NAME"))
	sb.WriteString(" ")
	sb.WriteString("port=")
	sb.WriteString(os.Getenv("DB_PORT"))

	dsn := sb.String()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database")
	}
}
