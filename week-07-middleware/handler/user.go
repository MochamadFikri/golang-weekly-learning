package handler

import "github.com/gin-gonic/gin"

func GetProfile(c *gin.Context) {
    c.JSON(200, gin.H{
        "user": "Fikri",
        "role": "Admin",
    })
}