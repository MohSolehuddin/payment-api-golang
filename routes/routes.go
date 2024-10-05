package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Routing function untuk Gin
func Routing() {
	router := gin.Default()

	// Mendefinisikan route "/soleh" dan menambahkan handler
	router.GET("/soleh", func(c *gin.Context) {
		// Mengembalikan response berupa JSON
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Soleh!",
		})
	})

	// Menjalankan server di port 8080
	router.Run("localhost:8080")  // Jalankan server di localhost:8080
}
