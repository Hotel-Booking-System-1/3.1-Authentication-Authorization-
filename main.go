package main

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/mubarik-siraji/booking-system/Routes"
	"github.com/mubarik-siraji/booking-system/infra"
)

func main() {
	slog.Info("Starting application...")
	infra.InitEnv()

	config := infra.Configurations

	// create gin app
	r := gin.Default()

	// connect database
	slog.Info("Connecting to database...")
	db := infra.ConnectDB()
 Routes.RegisterRoutes(r, db)

	r.Run(fmt.Sprintf(":%s", config.Port))

}
