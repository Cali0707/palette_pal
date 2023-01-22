package main

import (
	cohere "github.com/cohere-ai/cohere-go"
	"github.com/gin-gonic/gin"
)

func main() {
	port := ":3000"

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	r.Run(port)

	co, err := cohere.CreateClient("<<apiKey>>")
}
