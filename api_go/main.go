package main

import (
	"github.com/DanielGregorini/go-api-gin/controller"
	"github.com/DanielGregorini/go-api-gin/middleware"
	"github.com/DanielGregorini/go-api-gin/service"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	videoService    service.VideoService       = service.NewVideoService()
	videoController controller.VideoController = controller.NewVideoController(videoService)

	authService service.AuthService = service.NewAuthService(os.Getenv("SECRET_KEY"))

	userService    service.UserService       = service.NewUserService()
	userController controller.UserController = controller.NewUserController(userService)
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

	server.GET("/users", func(context *gin.Context) {
		users := userController.FindAll()
		context.JSON(200, users)
	})

	server.POST("/users", func(context *gin.Context) {
		context.JSON(200, userController.Save(context))
	})

	server.POST("/login", func(context *gin.Context) {
		user, error := userController.Login(context)

		if error != nil {
			context.JSON(401, gin.H{"error": "Usuário ou senha incorretos"})
			return
		}

		token, error := authService.GenerateToken(user)

		if error != nil {
			context.JSON(500, gin.H{"error": error.Error()})
			return
		}

		context.JSON(200, gin.H{
			"token": token,
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
				"password": "",
			},
		})
	})

	server.Run(":8080")
}
