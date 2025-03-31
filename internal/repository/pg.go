package repository

import (
	"rest-news/internal/entity"
	"rest-news/pkg/postgres"
)

type Repository interface {
	UpdateNews(entity.News) error
	GetNewsList(int, int) ([]entity.News, error)
	SaveUser(entity.User) error
	GetUser(string) (*entity.User, error)
}

type NewsRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *NewsRepo {
	return &NewsRepo{pg}
}
