package main

import (
	"fmt"
	"log"
	"os"

	"gitgub.com/gofiber/fiber/v2"
	"gitgub.com/joho/godotenv"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func main() {

	app := fiber.New()

	err := godotenv.Load(".env")

	if err != nil {

		log.Fatal("error loading env file")
	}

	dsn := os.Getenv("DB_DSN")

	dbClient, err = gorm.open(mysql.open(dsn), &gorm.Config{})

	sqlDB, _ := dbClient.DB()
	defer sqLDB.Close()

	if err != nil {

		log.Fatal("error connecting to database")

	}

	log.Println("Database connected")

	app.Get("/", func(c *fiber.Ctx) error {

		return c.status(Fiber.StatusOK).JSON(&fiber.Map{
			"message": "Everything is working fine",
		})
	}

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
