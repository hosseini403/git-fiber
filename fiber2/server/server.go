package server

import (
	"net/http"

import
"net/http"
	"fiberProject/Fiber2/middleware"

"github.com/gofiber/fiber/v2"
"github.com/Atarod61/middleware-go-fiber1/middleware"
	"github.com/gofiber/fiber/v2"
)

func Run(){
//Initialize Fiber
func Run() {
	// Initialize Fiber
	app := fiber.New()

	//Global Middleware
	app.use(middleware.RequestLogger)
	// Global Middlewares
	app.Use(middleware.GlobalLogger)

//Define Routes.
	app.Get("/",func(ctx*fiber.ctx) error{
		return ctx.status(http.statusok).Jason(fiber.Map{
			"message": "hello I am working",
		   )} 
		)}
		app.Get("/health",func (ctx*fiber.ctx) error {
			
			return ctx.status(http.statusok).Jason(fiber.Map{
				"healthy": true,
			})
	// Define Routes
	app.Get("/request-logger", middleware.RequestLogger, func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Request log is printed out",
		})
	})
	app.Get("/time-logger", middleware.TimeLogger, func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Time log is printed out",
		})
		app.Listen(":3000")
	})

}
	app.Listen(":3003")
}