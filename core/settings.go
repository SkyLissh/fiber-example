package core

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type settings struct {
	host string
	post string
	password string
	user string
	dbname string
	sslmode string

	ApiVersion string
	ProjectName string
}

func (s *settings) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		s.host, s.post, s.user, s.password, s.dbname, s.sslmode,
	)
}

func newSettings() *settings {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &settings{
		host: os.Getenv("DB_HOST"),
		post: os.Getenv("DB_PORT"),
		password: os.Getenv("DB_PASSWORD"),
		user: os.Getenv("DB_USER"),
		dbname: os.Getenv("DB_NAME"),
		sslmode: os.Getenv("DB_SSLMODE"),

		ApiVersion: os.Getenv("API_VERSION"),
		ProjectName: os.Getenv("PROJECT_NAME"),
	}
}

var Settings = newSettings()
