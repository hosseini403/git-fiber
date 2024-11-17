package main

import (
	"os"

	"github.com/gofiber/fiber"
)

func main() {

	app := fiber.New()

	err := godotenv.Load(".env")

	if err != nil {
		loge.fatal("error loading env file")

	}

	dsn := os.Getenv("DB_DSN")

}
