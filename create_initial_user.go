package main

import (
    "log"
    "go-project-api/internal/model"
    "go-project-api/internal/repository"
)

func main() {
    // Connect to the database
    userRepo, err := repository.ConnectDatabase()
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    // Create the initial user
    initialUser := &model.User{
        Username: "Admin1",
        Password: "super_secret",
        Email:    "admin@example.com",
        Phone:    "5551234567",
        IsSuperUser: true,
    }

    err = repository.CreateUser(userRepo, initialUser)
    if err != nil {
        log.Fatal("Failed to create the initial user:", err)
    }

    log.Println("Initial user created successfully")
}