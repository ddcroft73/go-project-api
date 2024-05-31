// This module will handle anything to do with authorization. Login,
// get token... etc handlers... this is the first stop after a request is sent from the client.
package handler

import (
	"github.com/gin-gonic/gin"
	"go-project-api/internal/service"
	_ "go-project-api/internal/util"
	"net/http"
)

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

	// Invoke the Register method from the AuthService
	err := h.authService.Register(username, password, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
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

// USer Operations.

func (h *AuthHandler) GetUser_Handle(c *gin.Context) {

}

func (h *AuthHandler) UpdateUser_Handle(c *gin.Context) {

}

func (h *AuthHandler) DeleteUser_Handle(c *gin.Context) {

}

func (h *AuthHandler) GetAllUsers_Handle(c *gin.Context) {

}
