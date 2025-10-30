package repository

import (
	"go-api/model"

	"gorm.io/gorm"
)

type VideoGameRepository struct {
	conn *gorm.DB
}

func NewVideoGameRepository(conn *gorm.DB) VideoGameRepository {
	return VideoGameRepository{ conn: conn }
}

func (p *VideoGameRepository) GetVideoGames() ([]model.VideoGame, error) {
	var video_games []model.VideoGame

	result := p.conn.Find(&video_games)
	return video_games, result.Error
}

func (p *VideoGameRepository) CreateVideoGame(video_game model.VideoGame) error {
	result := p.conn.Create(&video_game)
	return result.Error
}

func (p *VideoGameRepository) DeleteVideoGameById(id uint) error {
	result := p.conn.Delete(&model.VideoGame{}, id)
	return result.Error
}