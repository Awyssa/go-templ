package handlers

import (
	"go-templ/models"
	"go-templ/templates"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	store *models.TodoStore
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		store: models.NewTodoStore(),
	}
}

func (h *TodoHandler) Index(c *gin.Context) {
	todos := h.store.GetAll()
	component := templates.Index(todos)
	component.Render(c.Request.Context(), c.Writer)
}

func (h *TodoHandler) AddTodo(c *gin.Context) {
	task := c.PostForm("task")
	if task == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	todo := h.store.Add(task)
	component := templates.TodoItem(todo)
	component.Render(c.Request.Context(), c.Writer)
}

func (h *TodoHandler) ToggleTodo(c *gin.Context) {
	id := c.Param("id")
	todo, exists := h.store.Get(id)
	if !exists {
		c.Status(http.StatusNotFound)
		return
	}

	todo, _ = h.store.Update(id, !todo.Completed)
	component := templates.TodoItem(todo)
	component.Render(c.Request.Context(), c.Writer)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if success := h.store.Delete(id); !success {
		c.Status(http.StatusNotFound)
		return
	}
	c.Status(http.StatusOK)
}
