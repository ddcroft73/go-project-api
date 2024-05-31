package main

import (
	"go-project-api/internal/handler/route"
	"go-project-api/internal/security"
	_ "go-project-api/internal/util" // acesss to Log Writer
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    
    err := security.LoadEnvVars()
	if err != nil {
        log.Fatal("Error loading config.env file.")
    }

    route.SetupRouter(r)

    r.Run(":8080")
}