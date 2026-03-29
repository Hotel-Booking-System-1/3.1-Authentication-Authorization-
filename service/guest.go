package service

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/mubarik-siraji/booking-system/constant"
	"github.com/mubarik-siraji/booking-system/dtos"
	"github.com/mubarik-siraji/booking-system/helpers"
	"github.com/mubarik-siraji/booking-system/models"
	"github.com/mubarik-siraji/booking-system/repository"
	"golang.org/x/crypto/bcrypt"
)

type GuestService struct {
	repo *repository.GuestRepo
}

// Constructor
func RegisterGuestService(repo *repository.GuestRepo) *GuestService {
	return &GuestService{
		repo: repo,
	}
}

// CreateGuest registers a new guest
func (svc *GuestService) CreateGuest(data *dtos.CreateGuestDto) (int, error) {
	email := strings.ToLower(data.EmailAddress)

	_, err := svc.repo.GetGuestByEmail(email)
	if err == nil {
		errStr := "user with that email already exists"
		slog.Error(errStr)
		return http.StatusConflict, errors.New(errStr)
	}

	slog.Info("Hashing user password")
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to hash password")
		return http.StatusInternalServerError, errors.New(constant.DefaultErrorMsg)
	}

	data.Password = string(hashBytes)

	slog.Info("Creating new guest")
	err = svc.repo.CreateGuest(models.Guest{
		FullName:     data.FullName,
		EmailAddress: email,
		Password:     data.Password,
	})

	if err != nil {
		slog.Error("failed to create guest", "error", err)
		return http.StatusInternalServerError, errors.New(constant.FailedToCreateGuest)
	}

	slog.Info("Guest successfully created")
	return http.StatusCreated, nil
}

// LoginGuest authenticates a guest
func (svc *GuestService) LoginGuest(data dtos.LoginGuestRequest) (response *dtos.LoginGuestResponse, StatusCode int, err error) {
	email := strings.ToLower(data.EmailAddress)

	Guest, err := svc.repo.GetGuestByEmail(email)
	if err != nil {
		slog.Error("invalid email")
		return nil, http.StatusUnauthorized, errors.New(constant.InvalidCredentials)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(Guest.Password), []byte(data.Password)); err != nil {
		slog.Error("invalid password")
		return nil, http.StatusUnauthorized, errors.New(constant.InvalidCredentials)
	}

	accessToken, err := helpers.GenerateJWT(Guest.Role, Guest.EmailAddress, time.Now().Add(15*time.Minute).Unix(), false)

	if err != nil {
		slog.Error("failed to generate access token")
		return nil, http.StatusInternalServerError, errors.New(constant.DefaultErrorMsg)
	}

	refreshToken, err := helpers.GenerateJWT(Guest.Role, Guest.EmailAddress, time.Now().Add(72*time.Hour).Unix(), false)
	if err != nil {
		slog.Error("failed to generate refresh token")
		return nil, http.StatusInternalServerError, errors.New(constant.DefaultErrorMsg)
	}

	response = &dtos.LoginGuestResponse{
		Guest:        Guest,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, http.StatusOK, nil
}

// WhoAmI returns guest info by email
func (svc *GuestService) WhoAmI(email string) (*models.Guest, error) {
	user, err := svc.repo.GetGuestByEmail(email)
	if err != nil {
		return nil, errors.New(constant.DefaultErrorMsg)
	}

	return &user, nil
}
