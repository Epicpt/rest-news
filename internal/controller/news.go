package controller

import (
	"rest-news/internal/entity"

	"github.com/gofiber/fiber/v2"
)

func (r *newsRoutes) getNewsList(c *fiber.Ctx) error {
	page := c.Locals("page").(int)
	limit := c.Locals("limit").(int)

	news, err := r.u.GetNewList(page, limit)
	if err != nil {
		r.l.Err(err).Msg("Error getting news list")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get news",
		})
	}

	r.l.
		Info().
		Interface("news", news).
		Msg("News list fetched successfully")

	return c.Status(fiber.StatusOK).JSON(entity.NewsResponse{
		Success: true,
		News:    news,
	})
}

func (r *newsRoutes) editNews(c *fiber.Ctx) error {
	id, err := c.ParamsInt("Id")
	if err != nil {
		r.l.Err(err).Int("id", id).Msg("invalid news ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid news ID"})
	}

	var request entity.News
	if err := c.BodyParser(&request); err != nil {
		r.l.Err(err).Msg("invalid request format")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request format"})
	}

	err = r.u.EditNews(request)
	if err != nil {
		r.l.Err(err).Interface("request", request).Msg("failed to update news")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update news",
		})
	}

	r.l.
		Info().
		Interface("request", request).Msg("New edit successfully")

	return c.SendStatus(fiber.StatusCreated)
}
