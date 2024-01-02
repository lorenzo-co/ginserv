package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lorenzo-co/ginserv/models"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	return r
}

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	r.Run()
}
