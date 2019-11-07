package services

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func LoadDB() *gorm.DB {

	var (
		dbType     string = os.Getenv("DB_TYPE")
		dbHost     string = os.Getenv("DB_HOST")
		dbPort     string = os.Getenv("DB_PORT")
		dbUser     string = os.Getenv("DB_USER")
		dbName     string = os.Getenv("DB_NAME")
		dbPassword string = os.Getenv("DB_PASSWORD")
	)

	fmt.Println("dbType", dbType)
	fmt.Println("dbHost", dbHost)
	fmt.Println("dbPort", dbPort)
	fmt.Println("dbUser", dbUser)
	fmt.Println("dbName", dbName)
	fmt.Println("dbPassword", dbPassword)

	db, err := gorm.Open(
		dbType,
		fmt.Sprintf(
			"host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
			dbHost, dbPort, dbUser, dbName, dbPassword,
		),
	)
	if err != nil {
		log.Fatal("Error when loading database: ", err)
	}
	return db
}
