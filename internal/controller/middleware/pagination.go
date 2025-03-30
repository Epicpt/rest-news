package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Pagination(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	limit := c.Query("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		limitInt = 10
	}
	c.Locals("page", pageInt)
	c.Locals("limit", limitInt)

	return c.Next()
}
