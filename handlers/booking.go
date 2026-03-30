package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mubarik-siraji/booking-system/dtos"
	"github.com/mubarik-siraji/booking-system/service"
)

type BookingHandler struct {
	BookingService service.BookingService
}
func RegisterBookingHandler(bookingService service.BookingService) *BookingHandler {
	return &BookingHandler{
		BookingService:bookingService,
	}
}

func NewBookingHandler(s service.BookingService) *BookingHandler {
	return &BookingHandler{BookingService: s}
}

// 1. Create Booking
func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req dtos.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uncorrected information: " + err.Error()})
		return
	}

	result, err := h.BookingService.CreateBooking(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Booking-ka waa lagu guuleystay",
		"data":    result,
	})
}

// 2. Get All Bookings
func (h *BookingHandler) GetBookings(c *gin.Context) {
	var params dtos.BookingFilterParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 10
	}

	result, err := h.BookingService.GetAllBookings(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unavailabel booking information"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// 3. Update Booking Dates
func (h *BookingHandler) UpdateBooking(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookin id must be a number"})
		return
	}

	var req dtos.UpdateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.BookingService.UpdateBooking(uint(id), req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Booking is updated",
		"data":    result,
	})
}

// 4. Update Status
func (h *BookingHandler) UpdateStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be anumber"})
		return
	}

	var req dtos.UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is not correct or sended"})
		return
	}

	if err := h.BookingService.UpdateStatus(uint(id), req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "(Status) is changed"})
}

// 5. Cancel Booking
func (h *BookingHandler) CancelBooking(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a number"})
		return
	}

	if err := h.BookingService.CancelBooking(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "booking is cancelled (Cancelled)"})
}