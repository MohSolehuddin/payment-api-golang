package main

import (
	"fmt"
	"net/http"

	"github.com/MohSolehuddin/payment-api-golang/services"
	"github.com/gin-gonic/gin"
)

func main() {
		filename := "history.json"
		services.CreateHistory(filename, "User logged in", nil)
	
		transfer := &services.ActivityTransfer{
			Activity: "Money Transfer",
			From:     "Alice",
			To:       "Bob",
			Nominal:  "1000",
		}
		services.CreateHistory(filename, "Transfer activity recorded", transfer)

	fmt.Println("Payment API is ready in localhost:8080")
	router := gin.Default()
	router.POST("/payment", func(c *gin.Context) {
		// Mengembalikan response berupa JSON
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Soleh!",
		})
	})
	router.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Soleh!",
		})
	})
	router.POST("/logout", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Soleh!",
		})
	})
	router.Run()
}