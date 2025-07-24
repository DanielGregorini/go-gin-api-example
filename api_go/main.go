package main

import (
	"github.com/DanielGregorini/go-api-gin/controller"
	"github.com/DanielGregorini/go-api-gin/service"
	"github.com/DanielGregorini/go-api-gin/middleware"
	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.NewVideoService()
	videoController controller.VideoController = controller.NewVideoController(videoService)

	userService service.UserService = service.NewUserService()
	userController controller.UserController = controller.NewUserController(userService)
)

func main() {
	server := gin.Default()

	server.Use(middleware.PrintLogger())

	server.GET("/videos", func(context *gin.Context) {
		videos := videoController.FindAll()
		context.JSON(200, videos)
	})

	server.POST("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})

	server.GET("/users", func(context *gin.Context) {
		users := userController.FindAll()
		context.JSON(200, users)
	})

	server.POST("/users", func(context *gin.Context) {
		context.JSON(200, userController.Save(context))
	})

	server.Run(":8080")
}