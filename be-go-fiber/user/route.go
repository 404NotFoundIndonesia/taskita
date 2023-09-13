package user

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func New(db *sql.DB, app fiber.Router) {
	r := userRepository{database: db}
	s := NewService(&r)

	app.Post("/register", registerHandler(s))
}

func registerHandler(s service) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var request RegisterRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "you have incomplete registration data",
			})
		}

		result, err := s.Register(c.Context(), request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "user registered in successfully",
			"data":    result,
		})
	}
}
