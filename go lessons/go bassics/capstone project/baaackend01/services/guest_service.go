package services

import (
	"apdirizakismail/baaackend01-dayone/models"
	"errors"
)

type GuestService struct {
	guests []models.Guest
}

func NewGuestService() *GuestService {
	return &GuestService{
		guests: []models.Guest{},
	}
}

func (s *GuestService) CreateGuest(guest *models.Guest) (int, error) {
	guest.ID = uint(len(s.guests) + 1)
	s.guests = append(s.guests, *guest)
	return 200, nil
}

func (s *GuestService) GetAllGuests() ([]models.Guest, error) {
	return s.guests, nil
}

func (s *GuestService) GetGuestByID(id uint) (*models.Guest, error) {
	for _, g := range s.guests {
		if g.ID == id {
			return &g, nil
		}
	}
	return nil, errors.New("Guest not found")
}

func (s *GuestService) UpdateGuest(id uint, updatedGuest *models.Guest) (int, error) {
	for i, g := range s.guests {
		if g.ID == id {
			updatedGuest.ID = id
			s.guests[i] = *updatedGuest
			return 200, nil
		}
	}
	return 404, errors.New("Guest not found")
}

func (s *GuestService) DeleteGuest(id uint) (int, error) {
	for i, g := range s.guests {
		if g.ID == id {
			s.guests = append(s.guests[:i], s.guests[i+1:]...)
			return 200, nil
		}
	}
	return 404, errors.New("Guest not found")
}

// package service

// import (
// 	"apdirizakismail/baaackend01-dayone/repository"
// 	"net/http"
// )

// type GuestService struct {
// 	Repo *repository.GuestRepo
// }

// func NewGuestService(repo *repository.GuestRepo) *GuestService {
// 	return &GuestService{Repo: repo}
// }

// func (s *GuestService) CreateGuest(data *models.Guest) (int, error) {
// 	err := s.Repo.CreateGuest(data)
// 	if err != nil {
// 		return http.StatusBadRequest, err
// 	}
// 	return http.StatusCreated, nil
// }

// func (s *GuestService) GetAllGuests() ([]models.Guest, error) {
// 	return s.Repo.GetAllGuests()
// }

// func (s *GuestService) GetGuestByID(id uint) (models.Guest, error) {
// 	return s.Repo.GetGuestByID(id)
// }

// func (s *GuestService) UpdateGuest(id uint, data *models.Guest) (int, error) {
// 	err := s.Repo.UpdateGuest(id, data)
// 	if err != nil {
// 		return http.StatusBadRequest, err
// 	}
// 	return http.StatusOK, nil
// }

// func (s *GuestService) DeleteGuest(id uint) (int, error) {
// 	err := s.Repo.DeleteGuest(id)
// 	if err != nil {
// 		return http.StatusBadRequest, err
// 	}
// 	return http.StatusOK, nil
// }
