package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/programaticreviews/golang-gin/service"
)

type VideoController interface {
	Findall() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	print("controller returned for video entity")
	return &controller{
		service: service,
	}
}

func (c *controller) Findall() []entity.Video {
	print("all videos are found in controller")
	return c.service.Findall()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	print(err)
	if err != nil {
		return err
	}
	c.service.Save(video)
	print("video saved in controller")
	return nil
}
