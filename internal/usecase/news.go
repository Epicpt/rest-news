package usecase

import (
	"rest-news/internal/entity"
	"rest-news/internal/repository"
)

type UseCase struct {
	repo repository.Repository
}

func NewUseCase(repo repository.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) EditNews(news entity.News) error {
	if err := uc.repo.UpdateNews(news); err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) GetNewList(page, limit int) ([]entity.News, error) {
	newsList, err := uc.repo.GetNewsList(page, limit)
	if err != nil {
		return nil, err
	}
	return newsList, nil
}
