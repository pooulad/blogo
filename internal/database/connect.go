package database

import (
	"fmt"

	"github.com/pooulad/blogo/internal/config"
	"github.com/pooulad/blogo/internal/database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB    *gorm.DB
	Model model.Models
}

func Connect(cfg *config.Config) (*Store, error) {
	var (
		host     = cfg.DB.Postgresql.Host
		username = cfg.DB.Postgresql.Username
		password = cfg.DB.Postgresql.Password
		dbname   = cfg.DB.Postgresql.DbName
		port     = cfg.DB.Postgresql.Port
		sslmode  = cfg.DB.Postgresql.SslMode
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, username, password, dbname, port, sslmode)
	fmt.Println("this is dsn:",dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{}, &model.Post{})
	if err != nil {
		return nil, err
	}

	model := model.NewModels()

	return &Store{
		DB:    db,
		Model: model,
	}, nil
}
