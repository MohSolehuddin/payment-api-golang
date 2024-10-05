package main

import (
	"fmt"

	"github.com/MohSolehuddin/payment-api-golang/routes"
)

func main() {
	fmt.Println("Payment API is ready in localhost:8080")
	routes.Routing()
}