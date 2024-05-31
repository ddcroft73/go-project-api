package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go-project-api/internal/model"
	_ "go-project-api/internal/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
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

	if len(expTime) == 0 {
		expTime = []int64{time.Now().Add(time.Hour * 24).Unix()}
	}
	// Define the token claims
	claims := jwt.MapClaims{
		"userID": user.ID,
		"exp":    expTime, // Token expiration time (e.g., 24 hours from now)
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	api_key := os.Getenv("API_KEY")
	secretKey := []byte(api_key)

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
