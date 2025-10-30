package db

import (
	"fmt"
	"go-api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "localhost"
	port = 5432
	user = "docker"
	password = "docker"
	db_name = "postgres"
)

func ConnectDB() (*gorm.DB, error) {
	psql_info := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, db_name,
	)

	db, err := gorm.Open(postgres.Open(psql_info), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	
	db.AutoMigrate(&model.VideoGame{})

	fmt.Println("[DB]: Conexão com o banco de dados realizada com sucesso!")
	return db, nil
} 