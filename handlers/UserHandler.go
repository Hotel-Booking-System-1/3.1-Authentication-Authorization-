package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mubarik-siraji/booking-system/dtos"
	"github.com/mubarik-siraji/booking-system/service"
)

type UserHandler struct {
	userService service.UserService
}

func RegisterUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var requestBody dtos.CreateUserDto

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "Invalid request body",
			"is_success": false,
		})
		return
	}

	statusCode, err := h.userService.CreateUser(&requestBody)
	if err != nil {
		slog.Error("failed to create user", "error", err.Error())
		c.JSON(statusCode, gin.H{
			"message":    err.Error(),
			"is_success": false,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "User created successfully",
		"is_success": true,
	})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var requestBody dtos.LoginUserRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "Invalid request body",
			"is_success": false,
		})
		return
	}
	response, statusCode, err := h.userService.LoginUser(requestBody)
	if err != nil {
		slog.Error("failed to login user", "error", err.Error())
		c.JSON(statusCode, gin.H{
			"message":    err.Error(),
			"is_success": false,
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"message":    "Login successful",
		"is_success": true,
		"data":       response,
	})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	role := c.Query("role")
	isActive := c.Query("is_active")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	response, statusCode, err := h.userService.GetAllUsers(page, limit, role, isActive)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"status":  "success",
		"message": "Users fetched successfully",
		"data":    response,
	})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid user ID",
			"data":    nil,
		})
		return
	}

	userEmail, _ := c.Get("userEmail")
	userRole, _ := c.Get("userRole")

	response, statusCode, err := h.userService.GetUserByID(uint(id), fmt.Sprintf("%v", userEmail), fmt.Sprintf("%v", userRole))
	if err != nil {
		c.JSON(statusCode, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"status":  "success",
		"message": "User fetched successfully",
		"data": gin.H{
			"user": response,
		},
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	message, statusCode, err := h.userService.Logout()
	if err != nil {
		c.JSON(statusCode, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"status":  "success",
		"message": message,
		"data":    nil,
	})
}
