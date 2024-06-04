package service

import (
	"errors"
	_ "fmt"
	"go-project-api/internal/model"
	"go-project-api/internal/repository"
	"go-project-api/internal/security"
	_ "go-project-api/internal/util"
	"log"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(username, password string) (string, error) {

	user, err := repository.FindByUsername(s.userRepo, username)
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

func (s *AuthService) Register(username, password, email, phone string) error {
	// ...

	exists, err := repository.UsernameExists(s.userRepo, username)
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
		Phone:    phone,
	}

	err = repository.CreateUser(s.userRepo, user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return err
	}

	return nil
}

func (s *AuthService) UpdateUserByID(userID int, updatedUser *model.User) (*model.User, error) {

	existingUser, err := repository.GetUserByID(s.userRepo, userID)
	if err != nil {
		return nil, err
	}

	if updatedUser.Username != "" {
		existingUser.Username = updatedUser.Username
	}
	if updatedUser.Password != "" {
		// need to hash the new password before saving.
		hashedPassword, err := security.HashPassword(updatedUser.Password)
		if err != nil {
			return nil, err
		}
		existingUser.Password = hashedPassword
	}
	if updatedUser.Email != "" {
		existingUser.Email = updatedUser.Email
	}
	if updatedUser.Phone != "" {
		existingUser.Phone = updatedUser.Phone
	}

	// Save the updated user to the database
	err = repository.UpdateUser(s.userRepo, existingUser)
	if err != nil {
		return nil, err
	}

	return existingUser, nil

}

func (s *AuthService) GetUser(userID int) (*model.User, error) {

	user, err := repository.GetUserByID(s.userRepo, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
