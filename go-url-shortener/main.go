package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HelloHandlerFunc (c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hey Go URL Shortener!",
	})
}

func main() {
	// Create default engine instance.
	r := gin.Default()
	r.GET("/", HelloHandlerFunc)

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error %v", err))
	}
}