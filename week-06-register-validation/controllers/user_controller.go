package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang-weekly/week-06-register-validation/models"
)

func Register(c *gin.Context) {
	var input models.RegisterInput

	// Bind data JSON ke struct RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data JSON tidak valid"})
		return
	}

	// Validasi manual
	if strings.TrimSpace(input.Name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama wajib diisi"})
		return
	}

	if !strings.Contains(input.Email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format email tidak valid"})
		return
	}

	if len(input.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password minimal harus 6 karakter"})
		return
	}

	// Jika semua validasi lolos, kirim respon sukses
	c.JSON(http.StatusOK, gin.H{
		"pesan": "Registrasi berhasil",
		"user":  input,
	})
}