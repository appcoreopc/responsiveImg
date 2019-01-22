package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {

	fmt.Println("lets burn!!")
	ConfigureRestServices()
}

func ConfigureRestServices() {

	router.GET("/images/:id", posting)
	router.POST("/images", posting)
	router.PUT("/images", posting)
	router.DELETE("/images/:id", posting)
	router.Run(":3000")
}

func posting(c *gin.Context) {

	imageId := c.Param("id")

	c.JSON(200, gin.H{
		"message": imageId,
	})
}
