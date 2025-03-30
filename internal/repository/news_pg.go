package repository

import (
	"context"
	"fmt"

	"rest-news/internal/entity"
	"rest-news/pkg/postgres"
)

type Repository interface {
	UpdateNews(entity.News) error
	GetNewsList(int, int) ([]entity.News, error)
}

type NewsRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *NewsRepo {
	return &NewsRepo{pg}
}

func (r *NewsRepo) UpdateNews(news entity.News) error {
	ctx := context.Background()

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	query := `
			UPDATE "News" 
			SET "Title" = COALESCE(NULLIF($2, ''), "Title"),
				"Content" = COALESCE(NULLIF($3, ''), "Content")
			WHERE "Id" = $1`

	_, err = tx.Exec(ctx, query, news.ID, news.Title, news.Content)
	if err != nil {
		return fmt.Errorf("error updating news: %w", err)
	}

	if len(news.Categories) > 0 {
		query := `DELETE FROM "NewsCategories" WHERE "NewsId" = $1;`
		_, err = tx.Exec(ctx, query, news.ID)
		if err != nil {
			return fmt.Errorf("error deleting old categories: %w", err)
		}

		for _, categoryID := range news.Categories {
			query := `INSERT INTO "NewsCategories" ("NewsId", "CategoryId") VALUES ($1, $2);`
			_, err = tx.Exec(ctx, query, news.ID, categoryID)
			if err != nil {
				return fmt.Errorf("error inserting new categories: %w", err)
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (r *NewsRepo) GetNewsList(page, limit int) ([]entity.News, error) {
	ctx := context.Background()
	offset := (page - 1) * limit

	query := `
		SELECT n."Id", n."Title", n."Content", COALESCE(json_agg(nc."CategoryId") FILTER (WHERE nc."CategoryId" IS NOT NULL), '[]') AS categories
        FROM "News" n
        LEFT JOIN "NewsCategories" nc ON n."Id" = nc."NewsId"
        GROUP BY n."Id"
        ORDER BY n."Id" ASC
		LIMIT $1
		OFFSET $2;`

	rows, err := r.Pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch news: %w", err)
	}
	defer rows.Close()

	var newsList []entity.News

	for rows.Next() {
		var news entity.News

		err := rows.Scan(&news.ID, &news.Title, &news.Content, &news.Categories)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		newsList = append(newsList, news)
	}
	return newsList, nil
}
