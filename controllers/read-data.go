package controllers

import (
	"enity/zhoskiy-bench-go/gorm"
	"enity/zhoskiy-bench-go/models"

	"github.com/gofiber/fiber/v2"
)

func ReadData(c *fiber.Ctx) error {
	var bears []models.Bear

	gorm.Db.Find(&bears)

	return c.JSON(bears)
}
