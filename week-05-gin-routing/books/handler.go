package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, BookList)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, b := range BookList {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Buku tidak ditemukan"})
}

func CreateBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Format data tidak valid"})
		return
	}
	BookList = append(BookList, newBook)
	c.JSON(http.StatusCreated, gin.H{"message": "Buku berhasil ditambahkan", "data": newBook})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Format data tidak valid"})
		return
	}
	for i, b := range BookList {
		if b.ID == id {
			BookList[i] = updatedBook
			c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil diperbarui", "data": updatedBook})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Buku tidak ditemukan"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, b := range BookList {
		if b.ID == id {
			BookList = append(BookList[:i], BookList[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Buku tidak ditemukan"})
}
