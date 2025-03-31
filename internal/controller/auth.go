package controller

import (
	"errors"
	"rest-news/internal/entity"
	"rest-news/internal/services"
	"rest-news/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func (r *newsRoutes) register(c *fiber.Ctx) error {
	var user entity.User

	if err := c.BodyParser(&user); err != nil {
		r.l.Err(err).Msg("Cannot parse request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
		})
	}

	if err := validateUserInput(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hashedPass, err := services.HashPassword(user.Password)
	if err != nil {
		r.l.Err(err).Msg("Failed to hash password")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	user.Password = hashedPass
	if err = r.u.Create(user); err != nil {
		r.l.Err(err).Msg("Failed to create user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	r.l.
		Info().
		Interface("user", user).
		Msg("User save successfully")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func (r *newsRoutes) login(c *fiber.Ctx) error {
	var req entity.User

	if err := c.BodyParser(&req); err != nil {
		r.l.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	user, err := r.u.GetUser(req.Username)
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			r.l.Err(err).Msg("Invalid username or password")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid username or password",
			})
		}
		r.l.Err(err).Msg("failed to get user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get user",
		})
	}

	if !services.CheckPassword(req.Password, user.Password) {
		r.l.Err(err).Msg("Password is incorrect, hash not equal password")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	token, err := services.GenerateJWT(user.ID)
	if err != nil {
		r.l.Err(err).Msg("failed to generate token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}

	r.l.Info().Interface("user", user).Str("token", token).Msg("User logged in successfully")

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func validateUserInput(user entity.User) error {
	if len(user.Username) < 1 {
		return errors.New("login must be at least 1 characters long")
	}
	if len(user.Password) < 1 {
		return errors.New("password must be at least 1 characters long")
	}
	return nil
}
