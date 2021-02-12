// API Examples

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	// router.POST("/user/:name/*action", func(c *gin.Context) {
	// 	c.FullPath() == "/user/:name/*action" // true
	// })
	router.Run(":3000")
}

// router.GET("/someGet", getting)
// router.POST("/somePost", posting)
// router.PUT("/somePut", putting)
// router.DELETE("/someDelete", deleting)
// router.PATCH("/somePatch", patching)
// router.HEAD("/someHead", head)
// router.OPTIONS("/someOptions", options)

// By default it serves on :8080 unless a
// PORT environment variable was defined.

// router.Run()

// router.Run(":3000") for a hard coded port
