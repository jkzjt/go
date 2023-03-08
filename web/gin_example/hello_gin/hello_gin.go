package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a gin router with default middleware: logger and recovery (crash-free) middleware
	r := gin.Default()

	r.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin",
		})
	})

	// By default it serves on :8080 unless a PORT environment variable was defined.
	r.Run() // equals r.Run(":8080")
}
