package controller

import (
	"net/http"
	"strconv"
	
	"github.com/DanielGregorini/go-api-gin/entity"
	"github.com/DanielGregorini/go-api-gin/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll(context *gin.Context)
	Save(context *gin.Context)
	FindByID(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	Login(context *gin.Context)
}

type userController struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserController(userSvc service.UserService, authSrc service.AuthService) UserController {
	return &userController{userService: userSvc, authService: authSrc}
}

func (ctl *userController) FindAll(context *gin.Context) {
	users, err := ctl.userService.FindAll()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}

func (ctl *userController) FindByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := ctl.userService.FindByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}

func (ctl *userController) Save(context *gin.Context) {
	var user entity.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := ctl.userService.Save(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, created)
}

func (ctl *userController) Update(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var user entity.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = id

	updated, err := ctl.userService.Update(user)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, updated)
}

func (ctl *userController) Delete(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ctl.userService.Delete(id); err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.Status(http.StatusNoContent)
}

func (ctl *userController) Login(context *gin.Context) {
	var req entity.LoginRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctl.userService.Login(req)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := ctl.authService.GenerateToken(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}
