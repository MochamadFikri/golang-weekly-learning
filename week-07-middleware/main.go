package main

import (
    "github.com/gin-gonic/gin"
    "github.com/MochamadFikri/golang-weekly-learning/week-07-middleware/middleware"
    "github.com/MochamadFikri/golang-weekly-learning/week-07-middleware/handler"
)

func main() {
    r := gin.New()
    r.Use(gin.Recovery())
    r.Use(middleware.Logger())

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Welcome to Week 7!"})
    })

    userGroup := r.Group("/user")
    userGroup.Use(middleware.AuthMiddleware())
    userGroup.GET("/profile", handler.GetProfile)

    r.Run(":8080")
}