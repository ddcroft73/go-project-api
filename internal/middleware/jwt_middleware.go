package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

func JWTAuthMiddleware(c *gin.Context) {
	// Get the JWT from the request headers
	authHeader := c.GetHeader("Authorization")

	// Check if the token is present and correctly formatted
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found or malformed"})
		c.Abort()
		return
	}

	// Remove the 'Bearer ' prefix to get the actual token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Verify and parse the JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token algorithm is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			log.Fatal("API key not configured")
		}
		return []byte(apiKey), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid token: %v", err)})
		c.Abort()
		return
	}

	// Extract the user ID from the token claims if possible
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		c.Abort()
		return
	}

	// Store the user ID in the request context for further use
	c.Set("userID", int(userID))

	c.Next()
}
