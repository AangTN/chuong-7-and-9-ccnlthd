package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tenuser/myapp/handlers"
	"github.com/tenuser/myapp/repositories"
	"github.com/tenuser/myapp/services"
)

func main() {
	repo := repositories.NewInMemoryTaskRepository()
	service := services.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	router := gin.Default()
	api := router.Group("/api")
	api.GET("/tasks", handler.ListTasks)
	api.POST("/tasks", handler.CreateTask)
	api.GET("/tasks/:id", handler.GetTask)
	api.PUT("/tasks/:id", handler.UpdateTask)
	api.DELETE("/tasks/:id", handler.DeleteTask)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
