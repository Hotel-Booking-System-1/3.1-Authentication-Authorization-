package handlers

import (
	"net/http"
	"strconv"

	"github.com/mubarik-siraji/booking-system/dtos"
	"github.com/mubarik-siraji/booking-system/service"

	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	roomService service.RoomService
}
func RegisterROOMHandler(roomService service.RoomService) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
	}
}


func NewRoomHandler(svc service.RoomService) *RoomHandler {
	return &RoomHandler{roomService: svc}
}

// CreateRoom handles POST /rooms
func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var req dtos.CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.roomService.CreateRoom(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetRooms handles GET /rooms (with query params for filtering/pagination)
func (h *RoomHandler) GetRooms(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	filter := dtos.RoomFilter{
		Page:       page,
		Limit:      limit,
		Status:     c.Query("status"),
		RoomType:   c.Query("room_type"),
		RoomNumber: c.Query("room_number"),
	}

	res, err := h.roomService.GetRooms(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateRoom handles PATCH /rooms/:id
func (h *RoomHandler) UpdateRoom(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room id"})
		return
	}

	var req dtos.UpdateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.roomService.UpdateRoom(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteRoom handles DELETE /rooms/:id
func (h *RoomHandler) DeleteRoom(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room id"})
		return
	}

	if err := h.roomService.DeleteRoom(uint(id)); err != nil {
		// This handles the "Prevent deletion if active bookings exist" business rule error
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}