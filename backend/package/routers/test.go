package router

import (
	"inmes/package/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteTest(router fiber.Router) {
	router.Get("/test", handler.TestHandler)
}
