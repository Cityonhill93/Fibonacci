package main

import (
	api "Fibonacci/Api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("Numbers", api.GetNumbers)
	router.GET("Numbers/:index", api.GetNumber)

	router.Run("localhost:8080")
}
