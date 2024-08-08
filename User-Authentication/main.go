package main

import (
    "log"
    "github.com/labstack/echo/v4"
	"user-auth-api/storage"
	"user-auth-api/handlers"

)

func main() {
    initDB()
    defer DB.Close()

    e := echo.New()

    e.POST("/register", registerHandler)
    e.POST("/login", loginHandler)

    log.Println("Server starting on :8080")
    if err := e.Start(":8080"); err != nil {
        log.Fatalln(err)
    }
}
