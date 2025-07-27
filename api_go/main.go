package main

import (
	"os"

	"github.com/DanielGregorini/go-api-gin/config"
	"github.com/DanielGregorini/go-api-gin/controller"
	"github.com/DanielGregorini/go-api-gin/db"
	"github.com/DanielGregorini/go-api-gin/middleware"
	"github.com/DanielGregorini/go-api-gin/routes"
	"github.com/DanielGregorini/go-api-gin/service"
	"github.com/DanielGregorini/go-api-gin/entity"
	"github.com/gin-gonic/gin"
)

var (
	secretKey   string              = os.Getenv("SECRET_KEY")
	authService service.AuthService = service.NewAuthService(secretKey)

	cfg    = config.Load()
	dbConn = db.Connect(cfg)

	userService    service.UserService       = service.NewUserService(dbConn)
	userController controller.UserController = controller.NewUserController(userService, authService)

	videoService    service.VideoService       = service.NewVideoService(dbConn)
	videoController controller.VideoController = controller.NewVideoController(videoService, authService)
)

func main() {

	// migrations db
	dbConn.AutoMigrate(&entity.User{}, &entity.Video{})

	server := gin.Default()

	//middlewares
	server.Use(middleware.PrintLogger())
	server.Use(middleware.AuthMiddleware())

	//rota de teste
	server.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "api funcionando!"})
	})

	//rotas
	routes.UserRoute(server, userController)
	routes.VideoRoute(server, videoController)

	//inicia o servidor
	server.Run(":8080")
}
