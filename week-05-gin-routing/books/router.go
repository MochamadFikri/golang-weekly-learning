package books

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	bookGroup := r.Group("/buku")
	{
		bookGroup.GET("", GetAllBooks)
		bookGroup.GET("/:id", GetBookByID)
		bookGroup.POST("", CreateBook)
		bookGroup.PUT("/:id", UpdateBook)
		bookGroup.DELETE("/:id", DeleteBook)
	}
}
