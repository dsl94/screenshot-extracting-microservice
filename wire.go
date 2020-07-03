//+build wireinject

package main

import (
	"ffmpeg-exmple/controller"
	"ffmpeg-exmple/service"
	"github.com/google/wire"
)

func InitVideoController(service service.VideoService) controller.VideoController {
	wire.Build(controller.ProvideVideoController)

	return controller.VideoController{}
}