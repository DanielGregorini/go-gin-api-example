package routes

import (
	"github.com/DanielGregorini/go-api-gin/controller"
	"github.com/gin-gonic/gin"
)

func VideoRoute(router *gin.Engine, videoController controller.VideoController) {
	router.GET("/videos", videoController.FindAll)
	router.GET("/videos/:id", videoController.FindByID)
	router.POST("/videos", videoController.Save)
	router.PUT("/videos/:id", videoController.Update)
	router.DELETE("/videos/:id", videoController.Delete)
}