package main

import (
	"fmt"
	"inmes/config"
	router "inmes/package/routers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load configuration
	config.Load()

	env := os.Getenv("ENVELOPMENT")
	fmt.Println("Envilopment:", env)

	app := fiber.New()

	// Add logging middleware
	app.Use(logger.New())

	router.RegisterRoutes(app)

	// Listen on a port
	app.Listen(":8080")
	fmt.Println("server starting listen and serve on 0.0.0.0:8080")
}
