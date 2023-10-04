package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "example text!!")
	})

	err := router.Run(":8080")
	if err != nil {
		panic("failed to start web server, error: " + err.Error())
	}
}
