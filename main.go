package main

import (
	"log"
	"os"

	"enity/zhoskiy-bench-go/controllers"
	"enity/zhoskiy-bench-go/gorm"
	"enity/zhoskiy-bench-go/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	envMode := os.Getenv("APP_ENV")

	if envMode == "development" || envMode == "" {
		godotenv.Load()
	}

	port := os.Getenv("APP_PORT")
	app := fiber.New()

	gorm.Init()
	validator.Init()

	setupRoutes(app)

	log.Fatal(app.Listen(":" + port))
}

func setupRoutes(app *fiber.App) {
	app.Get("api/plaintext", controllers.Plaintext)
	app.Get("api/read-data", controllers.ReadData)
	app.Post("api/create-data", controllers.CreateData)
}
