package main

import (
	"github.com/gin-gonic/gin"
	"go-project-api/internal/handler"
	"go-project-api/internal/middleware"
	"go-project-api/internal/repository"
	"go-project-api/internal/security"
	"go-project-api/internal/service"
	_ "go-project-api/internal/util" // acesss to Log Writer
	"log"
)

func main() {
	r := gin.Default()

	err := security.LoadEnvVars()
	if err != nil {
		log.Fatal("Error loading config.env file.")
	}

	userRepo, err := repository.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	authMiddleware := middleware.NewAuthMiddleware(userRepo)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	r.POST("/login-access-token", authHandler.Login)
	r.POST("/register", authHandler.Register) //  CreateUser

	r.GET("/test-token", authHandler.TestToken)
	r.POST("/verify-email", authHandler.VerifyEmail)
	r.POST("/resend-email-verification", authHandler.ResendEmailVerification)
	r.POST("/reset-password", authHandler.ResetPassword)
	r.POST("/recover-password", authHandler.RecoverPassword)
	r.POST("/verify-2fa", authHandler.Verify2FA)
	r.POST("/resend-2fa", authHandler.Resend2FA)
	// User Operations
	r.GET("/get-user/:id", authHandler.GetUser_Handle)
	r.PUT("/update-user/:id", middleware.JWTAuthMiddleware, authMiddleware.AuthMiddlewareFunc, authHandler.UpdateUser_Handle) //

	r.DELETE("/delete-user/:id", authHandler.DeleteUser_Handle)
	r.GET("/get-all-users", authHandler.GetAllUsers_Handle)

	r.Use(authMiddleware.AuthMiddlewareFunc)
	r.Run(":8080")
}
