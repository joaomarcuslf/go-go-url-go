package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	configs "github.com/joaomarcuslf/go-go-url-go/configs"
	encoders "github.com/joaomarcuslf/go-go-url-go/encoders"
	store "github.com/joaomarcuslf/go-go-url-go/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context, serverConfig *configs.Server, optionsConfig *configs.Options) {
	var creationRequest UrlCreationRequest

	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := encoders.Encode(creationRequest.LongUrl)

	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	host := fmt.Sprintf("%s://%s/%s", optionsConfig.Schema, optionsConfig.Prefix, shortUrl)

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host,
	})
}

type CustomUrlCreationRequest struct {
	LongUrl   string `json:"long_url" binding:"required"`
	CustomUrl string `json:"custom_url" binding:"required"`
}

func CreateCustomUrl(c *gin.Context, serverConfig *configs.Server, optionsConfig *configs.Options) {
	var creationRequest CustomUrlCreationRequest

	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := creationRequest.CustomUrl

	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	host := fmt.Sprintf("%s://%s/%s", optionsConfig.Schema, optionsConfig.Prefix, shortUrl)

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	initialUrl, err := store.RetrieveInitialUrl(shortUrl)

	if initialUrl == "" || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "url not found"})
		return
	}

	// Check if URL contains ? and if so, add a & to the end of the URL

	if strings.Contains(initialUrl, "?") {
		initialUrl += "&"
	} else {
		initialUrl += "?"
	}

	c.Redirect(302, initialUrl+"source=go-go-url-go")
}
