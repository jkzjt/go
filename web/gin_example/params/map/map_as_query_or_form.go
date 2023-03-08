package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {
		args := c.QueryMap("args")
		filters := c.PostFormMap("filters")

		c.JSON(http.StatusOK, gin.H{
			"args":    args,
			"filters": filters,
		})
	})

	r.Run(":80")
}
