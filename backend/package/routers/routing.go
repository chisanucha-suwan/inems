package router

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CORSHandler returns a gin.HandlerFunc that handles CORS by adding the necessary headers to the response.
func CORSHandler(app *fiber.App) {
	// see more: https://docs.gofiber.io/api/middleware/cors/
	method := strings.Join([]string{
		fiber.MethodGet,
		fiber.MethodPost,
		fiber.MethodHead,
		fiber.MethodPut,
		fiber.MethodDelete,
		fiber.MethodPatch,
		fiber.MethodOptions,
	}, ", ")
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Access-Token",
		AllowCredentials: true,
		AllowMethods:     method,
	}))

	// add middleware after CORS configuration
	app.Use(func(c *fiber.Ctx) {
		// Access request method using fiber.Ctx.Method()
		if c.Method() != fiber.MethodOptions {
			// Set content type using fiber.Ctx.Set("Content-Type", "application/json")
			c.Set("Content-Type", "application/json")
		}

		if c.Method() == fiber.MethodOptions {
			// Abort with status using fiber.Ctx.Status(204).Send()
			c.Status(204).Send(nil)
			return
		}

		c.Next()
	})
}

func RegisterRoutes(app *fiber.App) {
	// route: /api/test
	apiGroup := app.Group("/api")
	testGroup := apiGroup.Group("/test")
	// route: /api/test/test
	RouteTest(testGroup)
}
