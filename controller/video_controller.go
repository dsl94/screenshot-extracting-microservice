package controller

import (
	"ffmpeg-exmple/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoController struct {
	VideoService service.VideoService
}

func ProvideVideoController(s service.VideoService) VideoController {
	return VideoController{VideoService: s}
}

func (v *VideoController) HandleVideos(c *gin.Context) {
	form, _ := c.MultipartForm()

	v.VideoService.GetScreenShotsFromVideo(form)

	c.JSON(http.StatusOK, gin.H{"message": "Videos successfully submitted for sampling"})
}