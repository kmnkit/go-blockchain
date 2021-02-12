package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Query String Parameters are parsed using the existing underlying request object.
	// The request responds to a url matching: /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		middlename := c.Query("middlename") // Shortcut for c.Request.URL.Query().Get("lastname")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s %s", firstname, middlename, lastname)
	})
	router.Run(":3000")
}
