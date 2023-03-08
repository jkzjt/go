package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	r.GET("/login", func(c *gin.Context) {
		ip := c.DefaultQuery("ip", "127.0.0.1")
		token := c.Query("token") // equals Request.URL.Query().Get("token")

		c.String(http.StatusOK, "from: %s, token: %s", ip, token)
	})

	r.Run(":80")
}
