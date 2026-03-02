package service

import (
	"errors"
	"fmt"
	"math"

	"github.com/mubarik-siraji/booking-system/dtos"
	"github.com/mubarik-siraji/booking-system/models"
	"github.com/mubarik-siraji/booking-system/repository"
)

type RoomService interface {
	CreateRoom(req dtos.CreateRoomRequest) (*dtos.RoomResponse, error)
	GetRooms(filter dtos.RoomFilter) (*dtos.PaginatedRoomResponse, error)
	GetRoomByID(id uint) (*dtos.RoomResponse, error)
	UpdateRoom(id uint, req dtos.UpdateRoomRequest) (*dtos.RoomResponse, error)
	DeleteRoom(id uint) error
}

type roomService struct {
	repo repository.RoomRepository
}

// RegisterRoomService uses the naming convention you provided.
// We return the Interface (RoomService) to keep the code flexible.
func RegisterRoomService(repo repository.RoomRepository) RoomService {
	return &roomService{
		repo: repo,
	}
}

// NewRoomService is an alias often used for the same purpose.
// You can keep both or choose the one your team prefers.
func NewRoomService(repo repository.RoomRepository) RoomService {
	return &roomService{repo: repo}
}

// CreateRoom maps the request to the model and persists it
func (s *roomService) CreateRoom(req dtos.CreateRoomRequest) (*dtos.RoomResponse, error) {
	if req.Status != "Available" {
		return nil, fmt.Errorf("unavailable status: %s", req.Status)
	}
	if req.Status == "Maintenance" {
		return nil, fmt.Errorf("cannot create or book a room with status: %s", req.Status)
	}
	if req.Status == "Inactive" {
		return nil, fmt.Errorf("cannot create or book a room with status: %s", req.Status)
	}

	room := &models.Room{
		RoomNumber:    req.RoomNumber,
		RoomType:      req.RoomType,
		PricePerNight: req.PricePerNight,
		Status:        req.Status,
		Description:   req.Description,
	}

	if err := s.repo.Create(room); err != nil {
		return nil, err
	}

	return mapToResponse(room), nil
}

// GetRooms handles pagination logic and total page calculation
func (s *roomService) GetRooms(f dtos.RoomFilter) (*dtos.PaginatedRoomResponse, error) {
	// Set defaults if not provided
	if f.Page <= 0 {
		f.Page = 1
	}
	if f.Limit <= 0 {
		f.Limit = 10
	}

	rooms, total, err := s.repo.GetAll(f)
	if err != nil {
		return nil, err
	}

	var data []dtos.RoomResponse
	for _, r := range rooms {
		data = append(data, *mapToResponse(&r))
	}

	// Calculate total pages for frontend navigation
	totalPages := int(math.Ceil(float64(total) / float64(f.Limit)))

	return &dtos.PaginatedRoomResponse{
		Data:       data,
		Total:      total,
		Page:       f.Page,
		Limit:      f.Limit,
		TotalPages: totalPages,
	}, nil
}

// GetRoomByID fetches a single room by primary key
func (s *roomService) GetRoomByID(id uint) (*dtos.RoomResponse, error) {
	room, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("room not found")
	}
	return mapToResponse(room), nil
}

// UpdateRoom applies partial updates safely using pointers from the DTO
func (s *roomService) UpdateRoom(id uint, req dtos.UpdateRoomRequest) (*dtos.RoomResponse, error) {
	room, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("room not found")
	}

	// Only update fields that were actually provided in the request
	if req.PricePerNight != nil {
		room.PricePerNight = *req.PricePerNight
	}
	if req.Status != nil {
		room.Status = *req.Status
	}
	if req.Description != nil {
		room.Description = *req.Description
	}

	if err := s.repo.Update(room); err != nil {
		return nil, err
	}

	return mapToResponse(room), nil
}

// DeleteRoom delegates the deletion (and booking check) to the repository
func (s *roomService) DeleteRoom(id uint) error {
	return s.repo.Delete(id)
}

// Internal helper function to convert Model to DTO
func mapToResponse(r *models.Room) *dtos.RoomResponse {
	return &dtos.RoomResponse{
		ID:            r.ID,
		RoomNumber:    r.RoomNumber,
		RoomType:      r.RoomType,
		PricePerNight: r.PricePerNight,
		Status:        r.Status,
		Description:   r.Description,
		CreatedAt:     r.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     r.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
