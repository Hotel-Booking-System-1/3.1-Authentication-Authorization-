package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mubarik-siraji/booking-system/handlers"
	"github.com/mubarik-siraji/booking-system/middleware"
	"github.com/mubarik-siraji/booking-system/repository"
	"github.com/mubarik-siraji/booking-system/service"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	// dependency chain
	userRepo := repository.RegisterUserRepo(db)
	userService := service.RegisterUserService(userRepo)
	userHandler := handlers.RegisterUserHandler(*userService)

	api := r.Group("/api")
	user := api.Group("/user")
	{
		user.POST("/create", userHandler.CreateUser)
		user.POST("/login", userHandler.LoginUser)
		// user.POST("/logout", userHandler.Logout)
	}

	// Protected routes
	apiUsers := api.Group("/users")
	apiUsers.Use(middleware.AuthMiddleware())
	{
		// only admin can get all users
		apiUsers.GET("", middleware.AuthorizeRoles("admin"), userHandler.GetAllUsers)
		apiUsers.GET("/:id", userHandler.GetUserByID)
	}

	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/logout", userHandler.Logout)
	}
}
