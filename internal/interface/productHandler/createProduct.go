package producthandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	productdomain "github.com/sebajax/go-clean-architecture/internal/domain/productDomain"
	"github.com/sebajax/go-clean-architecture/pkg/validate"
)

// body request schema for creating a new product
type createProductSchema struct {
	Name     string  `json:"name" validate:"required,min=5"`
	Sku      string  `json:"sku" validate:"required,min=8"`
	Category string  `json:"category" validate:"required,min=5,oneof=Laptop SmartPhone Tablet Headphones Camera Television Other"`
	Price    float64 `json:"price" validate:"required"`
}

// endpoint for creating a new product using closure
func CreateProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get body request
		var body createProductSchema

		// assign the request body into the body schema
		err := c.BodyParser(&body)
		if err != nil {
			// map the error & response via the middleware
			log.Error(err)
			return err
		}

		// validate body using the schema
		schemaErr, err := validate.Validate(&body)
		if err != nil {
			log.Error(err, schemaErr)
			return schemaErr
		}

		// no schema errors then map body to domain
		p, err := productdomain.New(
			body.Name,
			body.Sku,
			body.Category,
			body.Price,
		)
		if err != nil {
			log.Error(err)
			return err
		}

		// execute the service

		// success service execution
	}
}
