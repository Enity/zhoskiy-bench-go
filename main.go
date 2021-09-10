package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	envMode := os.Getenv("APP_ENV")

	if envMode == "development" || envMode == "" {
		godotenv.Load()
	}

	port := os.Getenv("APP_PORT")
	app := fiber.New()

	dsn := "root:123321@tcp(127.0.0.1:3306)/zhoskiy-bench-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Fatal(app.Listen(":" + port))
}
