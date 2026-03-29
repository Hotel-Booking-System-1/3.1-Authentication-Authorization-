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

	// 🔍 Debug (si aad u hubiso values-ka)
	log.Println("DB_HOST:", config.DB_HOST)
	log.Println("DB_USER:", config.DB_USER)
	log.Println("DB_NAME:", config.DB_NAME)
	log.Println("DB_PORT:", config.DB_PORT)

	// ❗ muhiim: hubi password-ka inuusan empty ahayn
	if config.DB_PASSWORD == "" {
		log.Fatal("❌ DB_PASSWORD is empty")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.DB_HOST,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ failed to connect database:", err)
	}

	// ✅ Test connection (ping)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("❌ failed to get sqlDB:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("❌ database not reachable:", err)
	}

	// ✅ Auto migrate
	err = db.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.Guest{},
	)
	if err != nil {
		log.Fatal("❌ failed to migrate database:", err)
	}

	log.Println("✅ database connected successfully 🚀")
	return db
}

// package infra

// import (
// 	"fmt"
// 	"log"

// 	"github.com/mubarik-siraji/booking-system/models"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func ConnectDB() *gorm.DB {

// 	config := Configurations
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
// 		config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT)

// 	// dsn := "host=localhost user=postgres password=admin dbname=schooldb port=5432 sslmode=disable"

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("❌ failed to connect database:", err)
// 	}
// 	if err := db.AutoMigrate(

// 		models.User{},
// 		models.Room{},
// 		models.Guest{},
// 	); err != nil {
// 		log.Fatal("❌ failed to migrate database:", err)
// 	}

// 	log.Println("✅ database connected successfully")
// 	return db
// }
