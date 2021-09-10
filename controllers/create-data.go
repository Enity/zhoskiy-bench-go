package controllers

import (
	"enity/zhoskiy-bench-go/gorm"
	"enity/zhoskiy-bench-go/models"
	"enity/zhoskiy-bench-go/validator"

	"github.com/gofiber/fiber/v2"
)

func CreateData(c *fiber.Ctx) error {
	bear := new(models.Bear)

	if err := c.BodyParser(&bear); err != nil {
		return err
	}

	err := validator.Validator.Struct(bear)

	if err != nil {
		return fiber.ErrBadRequest
	}

	gorm.Db.Create(&bear)

	return c.JSON(bear)
}
