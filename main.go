package main

import (
	"fmt"
	"github/imadbg01/url-shortening-api/model"

	"github.com/gin-gonic/gin"
)

func main() {

	model.Setup()

	fmt.Println("Hello Go URL Shortener !🚀 ")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	err := r.Run(":9000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
