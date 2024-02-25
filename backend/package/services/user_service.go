package services

import (
	"inmes/package/repository"
	"log"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) GetUsers() ([]UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	UserResponses := []UserResponse{}
	for _, user := range users {
		UserResponse := UserResponse{
			UserID:      user.UserID,
			Username:    user.Username,
			FullName_TH: user.FullName_TH,
		}
		UserResponses = append(UserResponses, UserResponse)
	}

	return UserResponses, nil
}

func (s userService) GetUser(id int) (*UserResponse, error) {
	user, err := s.userRepo.GetById(id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	userResponse := UserResponse{
		UserID:      user.UserID,
		Username:    user.Username,
		FullName_TH: user.FullName_TH,
	}

	return &userResponse, nil
}
