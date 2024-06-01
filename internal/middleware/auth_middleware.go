package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)


func AuthMiddleware(c *gin.Context) {
    // Get the user ID from the authenticated user (assuming you have authentication in place)
    userID := c.GetInt("userID")

    // Get the user ID from the URL parameter
    paramUserID := c.Param("id")

    // Convert the parameter user ID to an integer
    paramUserIDInt, err := strconv.Atoi(paramUserID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        c.Abort()
        return
    }

    // Check if the authenticated user is a superuser or the owner of the data
    if userID != paramUserIDInt && !IsSuperUser(userID) {
        c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
        c.Abort()
        return
    }

    c.Next()
}

func IsSuperUser(userID int) bool {
    // Check if the user is a superuser based on your application's logic
    // Return true if the user is a superuser, false otherwise
    // You can implement this function based on your specific requirements
    // For example, you can check the user's role in the database or a configuration file
    // Placeholder implementation:
    return userID == 1 // Assuming user with ID 1 is a superuser
}