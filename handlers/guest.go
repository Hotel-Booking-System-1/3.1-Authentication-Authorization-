package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mubarik-siraji/booking-system/dtos"
	"gorm.io/gorm"

	"github.com/mubarik-siraji/booking-system/repository"
	"github.com/mubarik-siraji/booking-system/service"
)

// GuestHandlers handles guest routes
type GuestHandlers struct {
	userService *service.GuestService
}

// RegisterGuestHandlers initializes GuestHandlers
func RegisterGuestHandlers(db *gorm.DB) *GuestHandlers {
	guestRepo := repository.RegisterGuestRepo(db) // ✅ sax
	guestSvc := service.RegisterGuestService(guestRepo)

	return &GuestHandlers{
		userService: guestSvc, // ✅ sax
	}
}

// CreateGuest handles POST /Guests/register
func (h *GuestHandlers) CreateGuest(c *gin.Context) {
	var requestBody dtos.CreateGuestDto

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "failed to bind request body",
			"is_success": false,
			"error":      err.Error(),
		})
		return
	}

	user, err := h.userService.CreateGuest(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "failed to create guest",
			"is_success": false,
			"error":      err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "guest created successfully",
		"is_success": true,
		"data":       user,
	})
}

// LoginGuest handles POST /Guests/login
func (h *GuestHandlers) LoginGuest(c *gin.Context) {
	var requestBody dtos.LoginGuestRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "failed to bind request body",
			"is_success": false,
		})
		return
	}

	resp, statusCode, err := h.userService.LoginGuest(requestBody) // ✅ waa function body
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message":    err.Error(),
			"is_success": false,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "guest logged in successfully",
		"is_success": true,
		"data":       resp,
	})
}

// WhoAmI handles GET /Guests/whoami
func (h *GuestHandlers) WhoAmI(c *gin.Context) {
	email := c.GetString("email")

	user, err := h.userService.WhoAmI(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message":    err.Error(),
			"is_success": false,
			"data":       nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"is_success": true,
		"data":       user,
	})
}
