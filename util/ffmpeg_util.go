package util

import (
	"log"
	"os/exec"
	"path/filepath"
)

func ExctractScreenshoots(filename string) {
	path := filepath.FromSlash("generated-images/"+filename)
	cmd := exec.Command("ffmpeg", "-i", "public/"+filename, "-vf", "fps=1", path+"output%d.jpg")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
