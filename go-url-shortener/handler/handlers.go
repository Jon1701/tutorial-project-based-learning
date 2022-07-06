package handler

import (
	"net/http"

	"github.com/Jon1701/project-based-learning/go-url-shortener/shortener"
	"github.com/Jon1701/project-based-learning/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

// Request model definition.
type URLCreationRequest struct {
	LongURL string `json:"long_url" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}

func CreateShortURL(c *gin.Context) {
	// Unmarshal JSON.
	var creationRequest URLCreationRequest;
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return;
	}

	shortURL := shortener.GenerateShortLink(creationRequest.LongURL, creationRequest.UserID)
	store.SaveURLMapping(shortURL, creationRequest.LongURL, creationRequest.UserID)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message": "short url created successfully",
		"short_url": host + shortURL,
	})
}

func HandleShortURLRedirect(c *gin.Context) {
	shortURL := c.Param("shortURL");
	initialURL := store.RetrieveInitialURL(shortURL);
	c.Redirect(302, initialURL)
}