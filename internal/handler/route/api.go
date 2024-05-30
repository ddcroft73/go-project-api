package route

import (
    "github.com/gin-gonic/gin"
    "go-project-api/internal/handler"
    "go-project-api/internal/service"
    "go-project-api/internal/repository"
	"log"
)

func SetupRouter(r *gin.Engine) {

    userRepo, err := repository.ConnectDatabase()
    if err != nil {
		log.Fatal(err)
    }    

    // Create an instance of AuthService
    authService := service.NewAuthService(userRepo)

    // Create an instance of AuthHandler
    authHandler := handler.NewAuthHandler(authService)

    // Register the routes
    r.POST("/login", authHandler.Login)
    r.POST("/register", authHandler.Register)

}
