package main

import (
	"github.com/gin-gonic/gin"
	"golang-weekly/week-06-register-validation/routes"
)

func main() {
	r := gin.Default()

	routes.UserRoutes(r)

	r.Run()
}
