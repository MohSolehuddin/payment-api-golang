package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routing() {
	router := gin.Default()
	router.GET("/soleh", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Soleh!",
		})
	})

	router.Run("localhost:8080")
}
