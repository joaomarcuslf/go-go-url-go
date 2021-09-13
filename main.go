package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	configs "github.com/joaomarcuslf/go-url-shortener/configs"
)

func main() {
	configuration, err := configs.FromEnv()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	gin.SetMode(configuration.Options.Mode)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	err = r.Run(fmt.Sprintf(":%s", configuration.Server.Port))

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server\n    Error: %v", err))
	}

	fmt.Println("Hello Go URL Shortener !🚀")
}
