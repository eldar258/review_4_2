package app

import (
	"log"
	"studentgit.kata.academy/eldar/review_4_2/repository"
	"studentgit.kata.academy/eldar/review_4_2/storage"
	"studentgit.kata.academy/eldar/review_4_2/storage/mysql"
)

type App struct {
}

func (a *App) CreateRepository() (repository.Repositorer, error) {
	sqlx, err := mysql.Connect()
	if err != nil {
		log.Printf("app.CreateRepository().mysql.Connect() err: %s", err)
		return nil, err
	}

	dao := storage.New(sqlx)

	return repository.New(dao), nil
}
