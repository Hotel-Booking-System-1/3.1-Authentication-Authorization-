package infra

import (
	"fmt"
	"log"

	"github.com/mubarik-siraji/booking-system/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	config := Configurations
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT)

	// dsn := "host=localhost user=postgres password=admin dbname=schooldb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ failed to connect database:", err)
	}
	if err := db.AutoMigrate(

		models.User{},
		models.Room{},
	); err != nil {
		log.Fatal("❌ failed to migrate database:", err)
	}

	log.Println("✅ database connected successfully")
	return db
}
