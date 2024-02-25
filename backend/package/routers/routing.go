package router

import (
	"database/sql"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
)

// CORSHandler returns a gin.HandlerFunc that handles CORS by adding the necessary headers to the response.
func CORSHandler(app *fiber.App) fiber.Handler {
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
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Access-Token",
		AllowMethods: method,
	}))

	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Access-Token")
		c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Set("Content-Type", "application/json")
		if c.Method() == "OPTIONS" {
			return c.Status(fiber.StatusNoContent).Send(nil)
		}
		return c.Next()
	}
}

func RegisterRoutes(db *sql.DB) *fiber.App {
	app := fiber.New()
	// Add logging middleware
	app.Use(logger.New())

	// CORS middle ware
	CORSHandler(app)

	dbs := sqlx.NewDb(db, "sqlserver")

	// route: /api/test
	apiGroup := app.Group("/api")
	testGroup := apiGroup.Group("/test")
	// route: /api/test/test
	RouteTest(testGroup)

	// route: /api/users
	usersGroup := apiGroup.Group("/users")
	// route: /api/users/all
	RouteUsers(dbs, usersGroup)

	return app
}
