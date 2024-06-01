package route

import (
	"github.com/gin-gonic/gin"
	"go-project-api/internal/handler"
	"go-project-api/internal/repository"
	"go-project-api/internal/service"
	"log"
	"go-project-api/internal/middleware"

)

func SetupRouter(r *gin.Engine) {

	userRepo, err := repository.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	r.POST("/login-access-token", authHandler.Login)
	r.POST("/register", authHandler.Register) // same as CreateUser
	r.GET("/test-token", authHandler.TestToken)
	r.POST("/verify-email", authHandler.VerifyEmail)
	r.POST("/resend-email-verification", authHandler.ResendEmailVerification)
	r.POST("/reset-password", authHandler.ResetPassword)
	r.POST("/recover-password", authHandler.RecoverPassword)
	r.POST("/verify-2fa", authHandler.Verify2FA)
	r.POST("/resend-2fa", authHandler.Resend2FA)
	// User Operations
	r.GET("/get-user/:id", authHandler.GetUser_Handle)
	
	r.PUT("/update-user/:id", middleware.AuthMiddleware, authHandler.UpdateUser_Handle)

	r.DELETE("/delete-user/:id", authHandler.DeleteUser_Handle)
	r.GET("/get-all-users", authHandler.GetAllUsers_Handle)
}
