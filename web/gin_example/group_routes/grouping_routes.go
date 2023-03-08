package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	user := r.Group("/user", func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusTemporaryRedirect, "Please login first")
		}
	})

	{
		user.GET("/profile", func(c *gin.Context) {
			full_path := c.FullPath()
			c.JSON(http.StatusOK, gin.H{
				"method":  "GET",
				"route":   full_path,
				"message": "a concise profile",
			})
		})

		user.POST("/settings", func(c *gin.Context) {
			full_path := c.FullPath()
			c.JSON(http.StatusOK, gin.H{
				"method":  "POST",
				"route":   full_path,
				"message": "you just posted a hansome avatar",
			})
		})

		user.PUT("/submit", func(c *gin.Context) {
			full_path := c.FullPath()
			c.JSON(http.StatusOK, gin.H{
				"method":  "PUT",
				"route":   full_path,
				"message": "you just modified your password",
			})
		})

		user.DELETE("/cancle", func(c *gin.Context) {
			full_path := c.FullPath()
			c.JSON(http.StatusOK, gin.H{
				"method":  "DELETE",
				"route":   full_path,
				"message": "sorry, we'll do our best in the future",
			})
		})
	}

	item := r.Group("/item")

	{
		item.GET("/list", func(c *gin.Context) {
			keywords := c.QueryArray("keywords")

			c.JSON(http.StatusOK, gin.H{
				"method":   "GET",
				"route":    c.FullPath(),
				"keywords": keywords,
				"message":  "below are the item on your keywords",
			})
		})

		item.POST("/filter", func(c *gin.Context) {
			keywords := c.QueryArray("keywords")
			filters := c.PostFormMap("filters")

			c.JSON(http.StatusOK, gin.H{
				"method":   "POST",
				"route":    c.FullPath(),
				"keywords": keywords,
				"filters":  filters,
				"message":  "below are the item on your keywords ans filters",
			})
		})
	}

	r.Run(":80")
}
