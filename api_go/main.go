package main

import (
	"os"

	"github.com/DanielGregorini/go-api-gin/config"
	"github.com/DanielGregorini/go-api-gin/controller"
	"github.com/DanielGregorini/go-api-gin/db"
	"github.com/DanielGregorini/go-api-gin/middleware"
	"github.com/DanielGregorini/go-api-gin/routes"
	"github.com/DanielGregorini/go-api-gin/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.NewVideoService()
	videoController controller.VideoController = controller.NewVideoController(videoService)

	secretKey   string              = os.Getenv("SECRET_KEY")
	authService service.AuthService = service.NewAuthService(secretKey)

	cfg    = config.Load()
	dbConn = db.Connect(cfg)

	userService    service.UserService       = service.NewUserService(dbConn)
	userController controller.UserController = controller.NewUserController(userService, authService)
)

func main() {

	server := gin.Default()

	server.Use(middleware.PrintLogger())

	server.Use(middleware.AuthMiddleware())

	server.GET("/videos", func(context *gin.Context) {
		videos := videoController.FindAll()
		context.JSON(200, videos)
	})

	server.POST("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})

	routes.UserRoute(server, userController)

	server.Run(":8080")
}
