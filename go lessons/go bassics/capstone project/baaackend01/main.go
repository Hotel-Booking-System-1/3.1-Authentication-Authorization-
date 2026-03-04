package main

import (
	"apdirizakismail/baaackend01-dayone/handlers"
	"apdirizakismail/baaackend01-dayone/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	service := services.NewGuestService()
	handler := handlers.NewGuestHandler(service)

	r.POST("/guests", handler.CreateGuest)
	r.GET("/guests", handler.GetAllGuests)
	r.GET("/guests/:id", handler.GetGuestByID)
	r.PUT("/guests/:id", handler.UpdateGuest)
	r.DELETE("/guests/:id", handler.DeleteGuest)

	r.Run(":8081")
}
