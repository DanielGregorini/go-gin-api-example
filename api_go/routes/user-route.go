package routes

import (
	"github.com/DanielGregorini/go-api-gin/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine, userController controller.UserController) {
	router.GET("/users", userController.FindAll)
	router.GET("/users/:id", userController.FindByID)
	router.POST("/users", userController.Save)
	router.PUT("/users/:id", userController.Update)
	router.DELETE("/users/:id", userController.Delete)
	router.POST("/login", userController.Login)
}