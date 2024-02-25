package router

import (
	"inems/package/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouteUsers(db *sqlx.DB, router fiber.Router) {
	h := handler.NewUserHandler(db)
	router.Get("/", h.GetUsers)
}
