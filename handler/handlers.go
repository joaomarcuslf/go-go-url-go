package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	configs "github.com/joaomarcuslf/go-go-url-shortener/configs"
	encoders "github.com/joaomarcuslf/go-go-url-shortener/encoders"
	store "github.com/joaomarcuslf/go-go-url-shortener/store"
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

	host := ""

	if serverConfig.Port == "80" {
		host = fmt.Sprintf("%s://%s/%s", optionsConfig.Schema, optionsConfig.Prefix, shortUrl)
	} else {
		host = fmt.Sprintf("%s://%s:%s/%s", optionsConfig.Schema, optionsConfig.Prefix, serverConfig.Port, shortUrl)
	}

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl, _ := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
