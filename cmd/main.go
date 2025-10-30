package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db_conn, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	video_game_repository := repository.NewVideoGameRepository(db_conn)
	video_game_service := service.NewVideoGameService(video_game_repository)
	video_game_controller := controller.NewVideoGameController(video_game_service)

	server.GET("/ping", func(ctx* gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/video-games", video_game_controller.GetVideoGames)

	server.POST("/video-games", video_game_controller.CreateVideoGame)

	server.DELETE("/video-games/:id", video_game_controller.DeleteVideoGameById)

	server.Run(":8000")
}