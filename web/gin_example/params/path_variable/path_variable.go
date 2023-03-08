package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// This handler will match /page/2 but will not match /page/ or /page
	r.GET("/page/:current", func(c *gin.Context) {
		currnet := c.Param("current")
		c.String(http.StatusOK, "current page: %s", currnet)
	})

	// However, this one will match /item/go/ and also /item/go/discount
	// If no other routers match /item/go, it will redirect to /item/go/
	r.GET("/item/:title/*action", func(c *gin.Context) {
		title := c.Param("title")
		action := c.Param("action")
		message := title + " now " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	r.POST("/user/:token/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:token/*action" // true
		c.String(http.StatusOK, "%t", b)
	})

	// This handler will add a new router for /item/list.
	// Exact routes are resolved before param routes, regardless of the order they were defined.
	// Routes starting with /item/list are never interpreted as /page/:current/... routes
	r.GET("/item/list", func(c *gin.Context) {
		c.String(http.StatusOK, "the available list are [...]")
	})

	r.Run(":80")
}
