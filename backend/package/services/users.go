package services

import (
	"inems/package/models"
	"inems/package/repository"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserService interface {
	GetUsers() ([]models.UserResponse, error)
	GetUser(id int) (*models.UserResponse, error)
}

type service struct {
	Repo repository.UserRepository
}

func NewUserService(db *sqlx.DB) *service {
	return &service{
		Repo: repository.NewUserRepository(db),
	}
}

func (s *service) GetUsers() ([]models.UserResponse, error) {
	users, err := s.Repo.GetAll()
	if err != nil {
		log.Println("Error getting users:", err)
		return nil, err
	}

	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = models.UserResponse{
			UserID:      user.UserID,
			Username:    user.Username,
			FullName_TH: user.FullName_TH,
		}
	}

	return userResponses, nil
}

func (s *service) GetUser(id int) (*models.UserResponse, error) {
	user, err := s.Repo.GetById(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	userResponse := models.UserResponse{
		UserID:      user.UserID,
		Username:    user.Username,
		FullName_TH: user.FullName_TH,
	}

	return &userResponse, nil
}
