package main

import (
	"fmt"

	"github.com/Jon1701/project-based-learning/go-url-shortener/handler"
	"github.com/Jon1701/project-based-learning/go-url-shortener/store"
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

	r.POST("/create-short-url", func (c *gin.Context) {
		handler.CreateShortURL(c)
	})

	r.GET("/:shortURL", func (c *gin.Context) {
		handler.HandleShortURLRedirect(c)
	})

	store.InitializeStore();

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error %v", err))
	}
}