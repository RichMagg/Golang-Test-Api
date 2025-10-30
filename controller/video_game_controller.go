package controller

import (
	"go-api/model"
	"go-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoGameController struct {
	service service.VideoGameService
}

func NewVideoGameController(service service.VideoGameService) VideoGameController {
	return VideoGameController{ service: service }
}

func (p *VideoGameController) GetVideoGames(ctx *gin.Context) {
	video_games, err := p.service.GetVideoGames()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, video_games)
}

func (p *VideoGameController) CreateVideoGame(ctx *gin.Context) {
	var video_game model.VideoGame
	
	if err := ctx.ShouldBindJSON(&video_game); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err := p.service.CreateVideoGame(video_game)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (p *VideoGameController) DeleteVideoGameById(ctx *gin.Context) {
	id_param := ctx.Param("id")
	id64, err := strconv.ParseUint(id_param, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	id := uint(id64)

	err = p.service.DeleteVideoGameById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}