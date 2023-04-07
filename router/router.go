package router

import (
	"github.com/FabricioAsat/todo-fullstack/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterRequest(app *fiber.App) {

	// Get methods
	app.Get("/api/tasks/", controllers.GET_ReadAllTasks)
	app.Get("/api/tasks/:taskID", controllers.GET_ReadOneTask)

	// Post methods
	app.Post("/api/tasks/", controllers.POST_CreateTask)

	// Put methods
	app.Put("/api/tasks/:taskID", controllers.PUT_UpdateTask)

	// Delete methods
	app.Delete("/api/tasks/:taskID", controllers.DELETE_DeleteTask)

}
