package service

import (
	"go-api/model"
	"go-api/repository"
)

type VideoGameService struct {
	repository repository.VideoGameRepository
}

func NewVideoGameService(repository repository.VideoGameRepository) VideoGameService {
	return VideoGameService { repository: repository }
}

func (p *VideoGameService) GetVideoGames() ([]model.VideoGame, error) {
	return p.repository.GetVideoGames()
}

func (p *VideoGameService) CreateVideoGame(video_game model.VideoGame) error {
	return p.repository.CreateVideoGame(video_game)
}

func (p *VideoGameService) DeleteVideoGameById(id uint) error {
	return p.repository.DeleteVideoGameById(id)
}