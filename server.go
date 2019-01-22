package main

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {

	imgtarget, _ := imaging.Open("test")

	dstImage128 := imaging.Resize(imgtarget, 128, 128, imaging.Lanczos)

	imaging.Save(dstImage128, "test")

	fmt.Println("lets burn!!")
	ConfigureRestServices()
}

func ConfigureRestServices() {

	//router.Static("/images", "/images")
	router.GET("/images/:id", getImage)
	router.POST("/images", uploadImage)

	router.DELETE("/images/:id", posting)
	router.Run(":3000")
}

func posting(c *gin.Context) {

	imageId := c.Param("id")

	c.JSON(200, gin.H{
		"message": imageId,
	})
}

func uploadImage(c *gin.Context) {
	file, _ := c.FormFile("image")
	log.Println(file.Filename)
	c.SaveUploadedFile(file, "/upload")
}

func getImage(c *gin.Context) {

	imageId := c.Param("id")

	if pusher := c.Writer.Pusher(); pusher != nil {
		if err := pusher.Push("/images/"+imageId+"jpg", nil); err != nil {
			log.Println("fail to push")
		}
	}

	c.JSON(200, gin.H{
		"message": imageId,
	})
}
