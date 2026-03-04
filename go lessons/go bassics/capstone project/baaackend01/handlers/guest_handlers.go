package handlers

import (
	"apdirizakismail/baaackend01-dayone/models"
	"apdirizakismail/baaackend01-dayone/services"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GuestHandler struct {
	Service *services.GuestService
}

func NewGuestHandler(service *services.GuestService) *GuestHandler {
	return &GuestHandler{Service: service}
}

func (h *GuestHandler) CreateGuest(c *gin.Context) {
	var request models.Guest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"is_success": false, "message": err.Error()})
		return
	}

	status, err := h.Service.CreateGuest(&request)
	if err != nil {
		c.JSON(status, gin.H{"is_success": false, "message": err.Error()})
		return
	}

	c.JSON(status, gin.H{"is_success": true, "message": "Guest created successfully!"})
}

func (h *GuestHandler) GetAllGuests(c *gin.Context) {
	guests, err := h.Service.GetAllGuests()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"is_success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"is_success": true, "data": guests})
}

func (h *GuestHandler) GetGuestByID(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	guest, err := h.Service.GetGuestByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"is_success": false, "message": "Guest not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"is_success": true, "data": guest})
}

func (h *GuestHandler) UpdateGuest(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var request models.Guest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"is_success": false, "message": err.Error()})
		return
	}

	status, err := h.Service.UpdateGuest(uint(id), &request)
	if err != nil {
		c.JSON(status, gin.H{"is_success": false, "message": err.Error()})
		return
	}

	c.JSON(status, gin.H{"is_success": true, "message": "Guest updated successfully!"})
}

func (h *GuestHandler) DeleteGuest(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	status, err := h.Service.DeleteGuest(uint(id))
	if err != nil {
		c.JSON(status, gin.H{"is_success": false, "message": err.Error()})
		return
	}

	c.JSON(status, gin.H{"is_success": true, "message": "Guest deleted successfully!"})
}
