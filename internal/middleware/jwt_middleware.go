package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"os"
	
)

func JWTAuthMiddleware(c *gin.Context) {
    // Get the JWT from the request headers
    tokenString := c.GetHeader("Authorization")

    // Check if the token is present
    if tokenString == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
        c.Abort()
        return
    }

    // Verify and parse the JWT
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		api_key := os.Getenv("API_KEY")
        return []byte(api_key), nil
    })

    if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
        return
    }

    // Extract the user ID from the token claims
    claims := token.Claims.(jwt.MapClaims)
    userID := int(claims["user_id"].(float64))

    // Store the user ID in the request context for further use
    c.Set("userID", userID)

    c.Next()
}