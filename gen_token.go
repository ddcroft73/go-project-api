package main

import (
    "go-project-api/internal/util"
	"go-project-api/internal/security"
     "go-project-api/internal/model"
	 "log"
)

func main() {
    err := security.LoadEnvVars()
	if err != nil {
        log.Fatal("Error loading config.env file.")
    }
	
	user := &model.User{
        ID: 3,
        Username: "Admin1",
        Password: "super_secret",
        Email:    "admin@example.com",
        Phone:    "5551234567",
        IsSuperUser: true,
    }

	token, err := security.GenerateJWTToken(user)
	if err != nil {
		log.Fatal(err.Error())
	}
    
    // write the token so I can pick it up for testing. only good for few days..,
    util.WriteLog(token)
}