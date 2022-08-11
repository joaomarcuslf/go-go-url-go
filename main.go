package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	configs "github.com/joaomarcuslf/go-go-url-go/configs"
	handler "github.com/joaomarcuslf/go-go-url-go/handler"
	store "github.com/joaomarcuslf/go-go-url-go/store"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load("./.env")
	configuration, err := configs.FromEnv()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	gin.SetMode(configuration.Options.Mode)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c, &configuration.Options)
	})

	r.POST("/custom-url", func(c *gin.Context) {
		handler.CreateCustomUrl(c, &configuration.Options)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	r.PUT("/:shortUrl", func(c *gin.Context) {
		handler.UpdateCustomUrl(c, &configuration.Options)
	})

	_, err = store.InitializeStore(&configuration.Redis)

	if err != nil {
		panic(err)
	}

	fmt.Println("Go Go, URL Go Starting Now! 🚀")

	err = r.Run(fmt.Sprintf(":%s", configuration.Server.Port))

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server\n    Error: %v", err))
	}
}
