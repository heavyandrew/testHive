package services

import (
	"golang.org/x/crypto/bcrypt"
	"testHive/internal/models"
	"testHive/internal/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	user.Password = ""
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) UserAlreadyExists(user *models.User) (*models.User, error) {
	user, err := s.UserRepo.GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Authenticate(username, password string) (*models.User, error) {
	user, err := s.UserRepo.GetUserByUsername(username)
	if err != nil || user == nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil
}
