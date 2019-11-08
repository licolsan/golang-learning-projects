package services

import (
	"fmt"
	"licolsan-mysql/models"
	"os"

	"github.com/alexcesaro/log/stdlog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var logger = stdlog.GetFromFlags()

func LoadDB() *gorm.DB {

	var (
		dbType     string = os.Getenv("DB_TYPE")
		dbHost     string = os.Getenv("DB_HOST")
		dbUser     string = os.Getenv("DB_USER")
		dbName     string = os.Getenv("DB_NAME")
		dbPassword string = os.Getenv("DB_PASSWORD")
	)

	db, err := gorm.Open(
		dbType,
		fmt.Sprintf(
			"%v:%v@(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPassword, dbHost, dbName,
		),
	)

	if err != nil {
		logger.Error("Error when loading database: ", err)
	} else {
		logger.Info("Database loaded")
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Profile{})

	db.LogMode(true)

	return db
}
