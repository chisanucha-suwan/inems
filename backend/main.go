package main

import (
	"fmt"
	"inmes/config"
	router "inmes/package/routers"
	"inmes/utilities"
	"os"
)

func main() {
	// Load configuration
	config.Load()

	env := os.Getenv("ENVELOPMENT")
	fmt.Println("Envilopment:", env)

	// create database connection
	db, err := utilities.Open(
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
	)
	if err != nil {
		// handle error
		fmt.Println("Error opening database connection:", err)
		return
	}
	defer db.Close()

	app := router.RegisterRoutes(db)

	// Listen on a port
	app.Listen(":8080")
	fmt.Println("server starting listen and serve on 0.0.0.0:8080")
}
