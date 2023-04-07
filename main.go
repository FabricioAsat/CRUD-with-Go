package main

import (
	"log"

	"github.com/FabricioAsat/todo-fullstack/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func initEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}
}

func main() {

	initEnvVars()
	app := fiber.New()

	// HTTPRequest controlls
	router.RouterRequest(app)

	app.Listen(":3000")
}
