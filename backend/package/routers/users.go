package router

import (
	"inmes/package/handler"
	"inmes/package/repository"
	"inmes/package/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouteUsers(db *sqlx.DB, router fiber.Router) {
	// TODO:: package\routers\users.go:12:20: not enough arguments in call to repository.NewUserRepositoryDB
	// have ()
	// want (*sqlx.DB)
	userRepository := repository.NewUserRepositoryDB(db)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	router.Get("/all", userHandler.GetUsers)
}
