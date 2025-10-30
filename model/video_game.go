package model

type VideoGame struct {
	ID uint `gorm:"primaryKey"`
	Title string `gorm:"not null"`
	Platform string
	Release_Year int
	Price float64 `gorm:"not null"`
}

func (VideoGame) TableName() string {
	return "video_game"
}