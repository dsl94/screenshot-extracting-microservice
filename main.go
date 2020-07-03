package main

import (
	"ffmpeg-exmple/config"
	"ffmpeg-exmple/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	videoController := InitVideoController(service.ProvideVideoService())

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/image", "generated-images")
	router.POST("/video", videoController.HandleVideos)

	client := config.ConnectToEureka()
	// start client, register、heartbeat、refresh
	client.Start()

	err := router.Run()
	if err != nil {
		panic(err)
	}
}

