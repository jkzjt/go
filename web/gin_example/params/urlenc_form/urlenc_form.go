package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/form_post", func(c *gin.Context) {
		username := c.PostForm("username")
		role := c.DefaultPostForm("role", "user")

		c.JSON(http.StatusOK, gin.H{
			"status":   "posted",
			"username": username,
			"role":     role,
		})
	})

	r.Run(":80")
}
