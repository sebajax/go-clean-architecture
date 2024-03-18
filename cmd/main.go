package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/sebajax/go-clean-architecture/pkg/middleware"
)

func main() {
	// create fiber instance
	app := fiber.New()

	// add fiber middlewares
	app.Use(cors.New())
	app.Use(helmet.New())

	// add custom middlewares
	app.Use(middleware.ErrorHandler)

	// create health api endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Status ok - api running in port %s", os.Getenv("API_PORT")))
	})

	// create api group /api/{routes}
	// api := app.Group("/api")

	// create api group for product /api/product/{routes}
	// apiProduct := api.Group("/product")

	// fiber running
	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("API_PORT"))))
}
