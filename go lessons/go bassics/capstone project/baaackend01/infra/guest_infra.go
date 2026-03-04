package infra

import (
	"apdirizakismail/baaackend01-dayone/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Get configuration
	config := GetConfigurations()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate Guest table
	if err := db.AutoMigrate(&models.Guest{}); err != nil {
		log.Fatal("Failed to auto migrate database:", err)
	}

	DB = db
}

// package infra

// import (
// 	"apdirizakismail/baaackend01-dayone/models"
// 	"fmt"
// 	"log"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDB() {

// 	config := Configurations

// 	dsn := fmt.Sprintf(
// 		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
// 		config.DBHost,
// 		config.DBUser,
// 		config.DBPassword,
// 		config.DBName,
// 		config.DBPort,
// 	)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database")
// 	}

// 	// 🔥 Auto migrate User + Guest
// 	if err := db.AutoMigrate(

// 		&models.Guest{},
// 	); err != nil {
// 		log.Fatal("Failed to auto migrate database")
// 	}

// 	DB = db
// }
