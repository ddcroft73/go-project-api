package service

import (
	"errors"
	_ "fmt"
	"go-project-api/internal/model"
	"go-project-api/internal/repository"
	"go-project-api/internal/security"
	_ "go-project-api/internal/util"
)

//thisd file is where the work is actually done for the authenticATION. effectivly this is the next stop after a request is sent.

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(username, password string) (string, error) {

	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if !security.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := security.GenerateJWTToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Register(username, password, email string) error {
	// Implement the registration logic here
	// This might involve validating the input data, checking for duplicate usernames/emails,
	// creating a new user record in the database, etc.
	// Example:
	exists, err := s.userRepo.UsernameExists(username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username already exists")
	}
	user := &model.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	return s.userRepo.CreateUser(user)
}
