// This module will handle anything to do with authorization. Login,
// get token... etc handlers... this is the first stop after a request is sent from the client.
package handler

import (
	"github.com/gin-gonic/gin"
	"go-project-api/internal/service"
	"go-project-api/internal/model"
	_ "go-project-api/internal/util"
	"net/http"
    "strconv"
	"go-project-api/internal/middleware"
)

type UpdateUserRequest struct {
    Username string `json:"username" binding:"omitempty"`
    Password string `json:"password" binding:"omitempty"`
    Email    string `json:"email" binding:"omitempty"`
    Phone    string `json:"phone" binding:"omitempty"`
}

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	// Extract username and password from request
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Invoke the Login method from the AuthService
	token, err := h.authService.Login(username, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Register(c *gin.Context) {
	// Extract user information from request
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
    phone := c.PostForm("phone")

	// Invoke the Register method from the AuthService
	err := h.authService.Register(username, password, email, phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}


// USer Operations.
func (h *AuthHandler) GetUser_Handle(c *gin.Context) {

}

func (h *AuthHandler) UpdateUser_Handle(c *gin.Context) {
    // get the id of the current user
     authenticatedUserID := c.GetInt("userID")

     // Get the user ID from the URL parameter
     paramUserID := c.Param("id")
 
     // Convert the parameter user ID to an integer
     paramUserIDInt, err := strconv.Atoi(paramUserID)
     if err != nil {
         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
         return
     }
 
     // Check if the authenticated user is the owner of the data or a superuser
     if authenticatedUserID != paramUserIDInt && !middleware.IsSuperUser(authenticatedUserID) {
         c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
         return
     }

    // Bind the request payload to the UpdateUserRequest struct
    var updateReq UpdateUserRequest
    if err := c.ShouldBindJSON(&updateReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := &model.User{
        Username: updateReq.Username, 
        Password: updateReq.Password, 
        Email: updateReq.Email, 
        Phone: updateReq.Phone, 
    }

    updatedUser,err := h.authService.UpdateUserByID(paramUserIDInt, user)
    if err != nil {
        c.JSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
    }
    
    // return the updated user info, when all is kosherDill
    c.JSON(http.StatusOK, gin.H{"message": updatedUser}) 
}


func (h *AuthHandler) DeleteUser_Handle(c *gin.Context) {

}

func (h *AuthHandler) GetAllUsers_Handle(c *gin.Context) {

}



func (h *AuthHandler) TestToken(c *gin.Context) {

}

func (h *AuthHandler) VerifyEmail(c *gin.Context) {

}

func (h *AuthHandler) ResendEmailVerification(c *gin.Context) {

}

func (h *AuthHandler) ResetPassword(c *gin.Context) {

}

func (h *AuthHandler) RecoverPassword(c *gin.Context) {

}

func (h *AuthHandler) Verify2FA(c *gin.Context) {

}

func (h *AuthHandler) Resend2FA(c *gin.Context) {

}
