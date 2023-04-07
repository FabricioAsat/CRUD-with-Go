package controllers

import (
	"context"
	"time"

	"github.com/FabricioAsat/todo-fullstack/collection"
	"github.com/FabricioAsat/todo-fullstack/database"
	"github.com/FabricioAsat/todo-fullstack/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// POST
func POST_CreateTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var taskCollection = collection.GetCollection(DB, "Tasks")
	defer cancel()

	task := new(model.Task)

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	taskPayload := model.Task{
		Title:       task.Title,
		Description: task.Description,
		User:        task.User,
		IsDone:      false,
		CreatedAt:   time.Now(),
	}

	result, err := taskCollection.InsertOne(ctx, taskPayload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Creation error"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": result})
}

// Get
func GET_ReadOneTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var taskCollection = collection.GetCollection(DB, "Tasks")
	defer cancel()

	taskID := c.Params("taskID")
	var result model.Task

	objID, _ := primitive.ObjectIDFromHex(taskID)

	err := taskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&result)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

// Get
func GET_ReadAllTasks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var taskCollection = collection.GetCollection(DB, "Tasks")
	defer cancel()

	cursor, err := taskCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	var allTasks []bson.M

	if cursor.All(ctx, &allTasks); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": allTasks})
}

// Put
func PUT_UpdateTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var taskCollection = collection.GetCollection(DB, "Tasks")
	defer cancel()

	taskID := c.Params("taskID")
	var task model.Task

	objID, _ := primitive.ObjectIDFromHex(taskID)

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "No body, no update"})
	}

	if task.Title == "" || task.Description == "" || task.User == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Send a complete body"})
	}

	// Actualizar los datos.
	var update = bson.M{"$set": bson.M{"title": task.Title, "description": task.Description, "user": task.User, "isdone": task.IsDone, "updatedat": time.Now()}}
	var filter = bson.M{"_id": objID}
	result, err := taskCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	if result.MatchedCount < 1 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error de matched"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": result})
}

// Delete
func DELETE_DeleteTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var taskCollection = collection.GetCollection(DB, "Tasks")
	taskID := c.Params("taskID")
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(taskID)
	result, err := taskCollection.DeleteOne(ctx, bson.M{"_id": objID})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
	}

	if result.DeletedCount < 1 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "ID not found"})

	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": "Deleted successful"})
}
