package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when loading .env file! ", err)
	}
}
