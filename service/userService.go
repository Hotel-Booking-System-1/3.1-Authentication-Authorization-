package service

import (
	"errors"
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

type UserService struct {
	repo *repository.UserRepo
}

func RegisterUserService(repo *repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) CreateUser(data *dtos.CreateUserDto) (int, error) {
	email := strings.ToLower(data.EmailAddress)

	_, err := svc.repo.GetUserByEmail(email)
	if err == nil {
		return http.StatusConflict, errors.New("user with this email already exists")
	}
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error accured ")
	}

	data.Password = string(hashBytes)
	err = svc.repo.CreateUser(models.User{
		Name: data.FullName,
		EmailAddress: email,
		Password: data.Password,
		Role: data.Role,	
	})
	if err != nil {
		return http.StatusInternalServerError, errors.New("failed to create user")
	}

	return http.StatusCreated, nil
}

//LoginUser authenticates a user and returns a JWT token if successful

func (svc *UserService) LoginUser(
	data dtos.LoginUserRequest,
) (*dtos.LoginUserResponse, int, error) {

	email := strings.ToLower(data.EmailAddress)

	user, err := svc.repo.GetUserByEmail(email)
	if err != nil {
		return nil, http.StatusUnauthorized, errors.New(constant.InvalidCredentials)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return nil, http.StatusUnauthorized, errors.New(constant.InvalidCredentials)
	}
	accessToken, err := helpers.GenerateJWT(
		user.Role,
		user.EmailAddress,
		time.Now().Add(15*time.Minute).Unix(),
		false,
	)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New(constant.DefualtErrorMsg)
	}

	refreshToken, err := helpers.GenerateJWT(
		user.Role,
		user.EmailAddress,
		time.Now().Add(72*time.Hour).Unix(),
		true,
	)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New(constant.DefualtErrorMsg)
	}

	response := &dtos.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}

	return response, http.StatusOK, nil
}

func (svc *UserService) GetAllUsers(page, limit int, role, isActive string) (*dtos.PaginatedUsersResponse, int, error) {
	users, total, err := svc.repo.GetAllUsers(page, limit, role, isActive)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("failed to fetch users")
	}

	var userResponses []dtos.UserResponse
	if len(users) == 0 {
		userResponses = make([]dtos.UserResponse, 0)
	}

	for _, user := range users {
		userResponses = append(userResponses, dtos.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.EmailAddress,
			Role:      user.Role,
			IsActive:  user.Is_Active,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		})
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))
	if totalPages == 0 && total > 0 {
		totalPages = 1
	}

	return &dtos.PaginatedUsersResponse{
		Items:      userResponses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, http.StatusOK, nil
}

func (svc *UserService) GetUserByID(requestedID uint, requesterEmail string, requesterRole string) (*dtos.UserResponse, int, error) {
	user, err := svc.repo.GetUserByID(requestedID)
	if err != nil {
		return nil, http.StatusNotFound, errors.New("user not found")
	}

	if !strings.EqualFold(requesterRole, string(models.RoleAdmin)) {
		if !strings.EqualFold(user.EmailAddress, requesterEmail) {
			return nil, http.StatusForbidden, errors.New("Forbidden: you don't have permission to access this resource")
		}
	}

	return &dtos.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.EmailAddress,
		Role:      user.Role,
		IsActive:  user.Is_Active,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, http.StatusOK, nil
}

func (svc *UserService) Logout() (string, int, error) {
	return "User logged out successfully.", http.StatusOK, nil
}
