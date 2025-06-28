package routes

import (
	"github.com/gin-gonic/gin"
	"golang-weekly/week-06-register-validation/controllers"
)

// Fungsi untuk mengatur routing user
func UserRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
}
