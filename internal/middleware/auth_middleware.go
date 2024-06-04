package middleware

import (
	"go-project-api/internal/repository"
	_ "go-project-api/internal/service"
	"go-project-api/internal/util"
	_ "go-project-api/internal/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// this middleware is used to authenticate the user of an endpoint. It will only allow access to the current user
// or to a superuser.

type AuthMiddleware struct {
	UserRepo *repository.UserRepository
}

func NewAuthMiddleware(userRepo *repository.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{UserRepo: userRepo}
}

func (m *AuthMiddleware) AuthMiddlewareFunc(c *gin.Context) {
	// Get the user ID from the authenticated user... the token
	userID, exists := c.Get("userID")

	util.WriteLog("User of the tokens userID: ", userID)

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		c.Abort()
		return
	}

	paramUserID := c.Param("id")
	paramUserIDInt, err := strconv.Atoi(paramUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		c.Abort()
		return
	}

	util.WriteLog("userID of the query: ", paramUserID)
	// Check if the authenticated user is a superuser or the owner of the data
	if userID.(int) != paramUserIDInt && !m.IsSuperUser(userID.(int)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}

func (m *AuthMiddleware) IsSuperUser(userID int) bool {
	if m.UserRepo == nil {
		log.Println("User repository is nil")
		return false
	}

	util.WriteLog("in IsSuperUser. userId = ", userID)
	util.WriteLog("in IsSuperUser. m.UserRepo = ", m.UserRepo)

	user, err := repository.GetUserByID(m.UserRepo, userID)
	if err != nil || user == nil {
		log.Println("Error retrieving user or no user found:", err)
		return false
	}

	return user.IsSuperUser
}
