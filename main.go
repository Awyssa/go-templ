package main

import (
	"go-templ/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Initialize handlers
	todoHandler := handlers.NewTodoHandler()

	// Routes
	r.GET("/", todoHandler.Index)
	r.POST("/todo", todoHandler.AddTodo)
	r.POST("/todo/:id/toggle", todoHandler.ToggleTodo)
	r.DELETE("/todo/:id", todoHandler.DeleteTodo)

	log.Println("Server started on http://localhost:8080")
	r.Run(":8080")
}
