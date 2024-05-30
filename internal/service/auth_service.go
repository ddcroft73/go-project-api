package service

import (
    "errors"
    "go-project-api/internal/repository"
    "go-project-api/internal/model"
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
    // Implement the login logic here
    // This might involve checking the username and password against the database,
    // generating a JWT token upon successful authentication, etc.
    // Example:

	// ...
    user, err := s.userRepo.FindByUsername(username)
    if err != nil {
        return "", err
    }
    if user.Password != password {
        return "", errors.New("invalid credentials")
    }
    token, err := generateJWTToken(user)
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



func generateJWTToken(*model.User) (string, error){
    
	var token string
	return token, nil
}
