package security

import (
	"errors"
	"fmt"
	"go-project-api/internal/model"
	"go-project-api/internal/util"
	_ "go-project-api/internal/util"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func LoadEnvVars() error {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWTToken(user *model.User, expTime ...int64) (string, error) {

	// Set default expiration time if not provided
	var exp int64
	if len(expTime) == 0 {
		exp = time.Now().Add(time.Hour * 24).Unix() // 24 hours from now
	} else {
		exp = expTime[0] // Use the first element provided
	}

	// Define the token claims
	claims := jwt.MapClaims{
		"userID": user.ID,
		"exp":    exp, // Token expiration time
	}

	util.WriteLog("in GenerateJWTToken: user.ID in token = ", claims["userID"])

	// Create a new token with the specified claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Get the secret key from environment variables
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return "", errors.New("API_KEY environment variable not set")
	}
	secretKey := []byte(apiKey)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {

	api_key := os.Getenv("API_KEY")
	secretKey := []byte(api_key)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
