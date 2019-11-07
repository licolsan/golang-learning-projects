package main

import (
	"licolsan-postgres/initializers"
	// "licolsan-postgres/models"
	"licolsan-postgres/services"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	initializers.LoadEnvironment()
	db := services.LoadDB()
	defer db.Close()

	// db.AutoMigrate(&models.Production)
	// db.Create(&models.Product{Code: "L1212", Price: 1000})

	// var product models.Product

	// db.First(&product, 1)
	// db.First(&product, "code = ?", "L1212")
	// db.Model(&product).Update("Price", 2000)
	// db.Delete(&product)
}
