package controller

import (
	"rest-news/internal/controller/middleware"
	"rest-news/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type newsRoutes struct {
	u usecase.UseCase
	l zerolog.Logger
}

func NewNewsRoutes(router fiber.Router, u usecase.UseCase, l zerolog.Logger) {
	r := &newsRoutes{u, l}

	// Маршруты для авторизации и регистрации
	authGroup := router.Group("/auth")
	authGroup.Post("/register", r.register)
	authGroup.Post("/login", r.login)

	// Маршруты для новостей
	newsGroup := router.Group("/news", middleware.Auth)
	newsGroup.Get("/list", middleware.Pagination, r.getNewsList)
	newsGroup.Post("/edit/:Id", r.editNews)
}
