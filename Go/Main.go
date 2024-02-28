package main

import (
	api "Fibonacci/Api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("Numbers", api.GetNumbers)

	router.Run("localhost:8080")
}
