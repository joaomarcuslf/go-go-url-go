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

func formatUrl(optionsConfig *configs.Options, shortUrl string) string {
	return fmt.Sprintf("%s://%s/%s", optionsConfig.Schema, optionsConfig.Prefix, shortUrl)
}

func internalSaveFunction(shortUrl, longUrl string, optionsConfig *configs.Options) (string, error) {
	err := store.SaveUrlMapping(shortUrl, longUrl)

	if err != nil {
		return "", err
	}

	host := formatUrl(optionsConfig, shortUrl)

	return host, nil
}

type URLRequest struct {
	LongUrl   string `json:"long_url" binding:"required"`
	CustomUrl string `json:"custom_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context, optionsConfig *configs.Options) {
	var requestBody URLRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	host, err := internalSaveFunction(encoders.Encode(requestBody.LongUrl), requestBody.LongUrl, optionsConfig)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host,
	})
}

func CreateCustomUrl(c *gin.Context, optionsConfig *configs.Options) {
	var requestBody URLRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	host, err := internalSaveFunction(requestBody.CustomUrl, requestBody.LongUrl, optionsConfig)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":   "Custom url created successfully",
		"short_url": host,
	})
}

func UpdateCustomUrl(c *gin.Context, optionsConfig *configs.Options) {
	var requestBody URLRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := c.Param("shortUrl")
	longUrl := requestBody.LongUrl

	err := store.UpdateUrlMapping(shortUrl, longUrl)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	host := formatUrl(optionsConfig, shortUrl)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":   "Custom url created successfully",
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

	if strings.Contains(initialUrl, "?") {
		initialUrl += "&"
	} else {
		initialUrl += "?"
	}

	c.Redirect(302, initialUrl+"source=go-go-url-go")
}
