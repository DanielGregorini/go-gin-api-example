package controller

import (
	"github.com/DanielGregorini/go-api-gin/entity"
	"github.com/DanielGregorini/go-api-gin/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(context *gin.Context)  entity.Video
}

type videoController struct {
	service service.VideoService
}

func NewVideoController(service service.VideoService) VideoController {
	return &videoController{service: service}
}

func (c *videoController) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *videoController) Save(context *gin.Context) entity.Video {
	var video entity.Video

	if err := context.ShouldBindJSON(&video); err != nil {

		return entity.Video{}
	}

	saved := c.service.Save(video)
	return saved
}