package routes

import (
	"apdirizakismail/baaackend01-dayone/handlers"
	"apdirizakismail/baaackend01-dayone/middlewares"
	"apdirizakismail/baaackend01-dayone/services"

	"github.com/gin-gonic/gin"
)

func RegisterGuestRoutes(r *gin.Engine) {

	apiGroup := r.Group("/api")

	// // guestHandler := handlers.RegisterGuestHandlers()
	// guestHandler := handlers.RegisterGuestHandlers()
	service := services.NewGuestService()
	guestHandler := handlers.NewGuestHandler(service)

	guestsGroup := apiGroup.Group("/guests")
	{
		guestsGroup.POST("/", middlewares.Authenticate(), guestHandler.CreateGuest)
		guestsGroup.GET("/", middlewares.Authenticate(), guestHandler.GetAllGuests)
		guestsGroup.GET("/:id", middlewares.Authenticate(), guestHandler.GetGuestByID)
		guestsGroup.PUT("/:id", middlewares.Authenticate(), guestHandler.UpdateGuest)
		guestsGroup.DELETE("/:id", middlewares.Authenticate(), guestHandler.DeleteGuest)
	}
}
