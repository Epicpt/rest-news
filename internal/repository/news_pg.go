package repository

import (
	"rest-news/internal/entity"
)

type Repository interface {
	UpdateNews(entity.News) error
	GetNewsList() ([]entity.News, error)
}
