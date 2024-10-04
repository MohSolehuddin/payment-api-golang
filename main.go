package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Payment API is ready in localhost:8080")
	router := gin.Default()
	router.Run("localhost:8080")
}