package controller

import (
	"net/http"
	"strconv"

	"github.com/DanielGregorini/go-api-gin/entity"
	"github.com/DanielGregorini/go-api-gin/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll(context *gin.Context)
	Save(context *gin.Context)
	FindByID(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type videoController struct {
	videoService service.VideoService
	authService  service.AuthService
}

func NewVideoController(srcVideo service.VideoService, authService service.AuthService) VideoController {
	return &videoController{videoService: srcVideo, authService: authService}
}

func (ctl *videoController) FindAll(context *gin.Context) {
	users, err := ctl.videoService.FindAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, users)
}

func (ctl *videoController) FindByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := ctl.videoService.FindByID(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (ctl *videoController) Save(context *gin.Context) {
	var video entity.Video

	if err := context.ShouldBindJSON(&video); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := ctl.videoService.Save(video)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, created)
}

func (ctl *videoController) Update(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	oldVideo, err := ctl.videoService.FindByID(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	token := ctl.authService.ExtractTokenFromContext(context)


	if idInToken, err := ctl.authService.IsIDInToken(token, oldVideo.UserID); !idInToken || err != nil {
		context.JSON(http.StatusForbidden, gin.H{"error": "não autorizado"})
		return
	}

	var video entity.Video

	if err := context.ShouldBindJSON(&video); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	video.ID = oldVideo.ID

	updated, err := ctl.videoService.Update(video)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, updated)
}

func (ctl *videoController) Delete(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	video, err := ctl.videoService.FindByID(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	token := ctl.authService.ExtractTokenFromContext(context)

	if idInToken, err := ctl.authService.IsIDInToken(token, video.UserID); !idInToken || err != nil {
		context.JSON(http.StatusForbidden, gin.H{"error": "não autorizado"})
		return
	}

	if err := ctl.videoService.Delete(id); err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}