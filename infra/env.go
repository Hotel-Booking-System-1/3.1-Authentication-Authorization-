package infra


import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Appconfig struct {
	Port                  string
	DB_HOST               string
	DB_USER               string
	DB_PASSWORD           string
	DB_NAME               string
	DB_PORT               string
	AccsessJwtToKenSecret string
	RefreshJwtToKenSecret string
}

var Configurations Appconfig

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Configurations.Port = os.Getenv("PORT")
	Configurations.DB_HOST = os.Getenv("DB_HOST")
	Configurations.DB_USER = os.Getenv("DB_USER")
	Configurations.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	Configurations.DB_NAME = os.Getenv("DB_NAME")
	Configurations.DB_PORT = os.Getenv("DB_PORT")
	Configurations.AccsessJwtToKenSecret = os.Getenv("ACCSESS_TOKEN_JWT_SECRET")
	Configurations.RefreshJwtToKenSecret = os.Getenv("REFRESH_TOKEN_JWT_SECRET")
}
