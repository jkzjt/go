package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {
		keyword := c.Query("keyword")
		page := c.DefaultQuery("page", "1")

		sort := c.PostForm("sort")
		flag := c.DefaultPostForm("flag", "price DESC")

		c.JSON(http.StatusOK, gin.H{
			"keyword": keyword,
			"page":    page,
			"sort":    sort,
			"flag":    flag,
		})
	})

	r.Run(":80")
}
