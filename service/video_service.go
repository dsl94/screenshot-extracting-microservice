package service

import (
	"ffmpeg-exmple/util"
	"io"
	"log"
	"mime/multipart"
	"os"
)

type VideoService struct {

}

func ProvideVideoService() VideoService {
	return VideoService{}
}

func (s *VideoService) GetScreenShotsFromVideo(form *multipart.Form) {
	files := form.File["file"]

	for _, file := range files {
		open, _ := file.Open()

		filename := file.Filename
		out, err := os.Create("public/" + filename)
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(out, open)
		if err != nil {
			log.Fatal(err)
		}

		defer out.Close()

		go util.ExctractScreenshoots(filename)
	}
}
