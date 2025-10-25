package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tenuser/myapp/handlers"
	"github.com/tenuser/myapp/middlewares"
	"github.com/tenuser/myapp/repositories"
	"github.com/tenuser/myapp/services"
)

func main() {
	repo := repositories.NewInMemoryTaskRepository()
	service := services.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	router := gin.Default()

	// Global logging middleware (prints: [METHOD] URL - STATUS_CODE)
	router.Use(middlewares.LoggingMiddleware())

	api := router.Group("/api")
	api.GET("/tasks", handler.ListTasks)
	// protect creating, updating and deleting with a simple API key
	api.POST("/tasks", middlewares.SimpleAuthMiddleware(), handler.CreateTask)
	api.GET("/tasks/:id", handler.GetTask)
	api.PUT("/tasks/:id", middlewares.SimpleAuthMiddleware(), handler.UpdateTask)
	api.DELETE("/tasks/:id", middlewares.SimpleAuthMiddleware(), handler.DeleteTask)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
