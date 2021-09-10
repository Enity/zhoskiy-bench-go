package controllers

import "github.com/gofiber/fiber/v2"

func Plaintext(c *fiber.Ctx) error {
	return c.SendString("ok lets go")
}
