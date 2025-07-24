package controller

import (
	"github.com/DanielGregorini/go-api-gin/entity"
	"github.com/DanielGregorini/go-api-gin/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll() []entity.User
	Save(context *gin.Context) entity.User
	Login (context *gin.Context) (entity.User, error)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{service: service}
}

func (controller *userController) FindAll() []entity.User {
	return controller.service.FindAll()
}

func (controller *userController) Save(context *gin.Context) entity.User {
	var user entity.User

	if err := context.ShouldBindJSON(&user); err != nil {
		return entity.User{}
	}

	saved := controller.service.Save(user)
	return saved
}

func (controller *userController) Login(context *gin.Context) (entity.User, error) {

	var user entity.User

	if err := context.ShouldBindJSON(&user); err != nil {
		return entity.User{}, err
	}

	username := user.Username
	password := user.Password

	return controller.service.Login(username, password)
}