package main

import (
    "github.com/gin-gonic/gin"
    "go-project-api/internal/handler/route"
)

func main() {
	// setup rizouter
    r := gin.Default()

	// endpoints	
    route.SetupRouter(r)

    r.Run(":8080")
}