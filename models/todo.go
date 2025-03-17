package models

import (
	"fmt"
	"sync"
	"time"
)

type Todo struct {
	ID        string    `json:"id"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type TodoStore struct {
	sync.RWMutex
	todos map[string]*Todo
}

func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos: make(map[string]*Todo),
	}
}

func (s *TodoStore) Add(task string) *Todo {
	s.Lock()
	defer s.Unlock()

	id := fmt.Sprintf("%d", time.Now().UnixNano())
	todo := &Todo{
		ID:        id,
		Task:      task,
		Completed: false,
		CreatedAt: time.Now(),
	}
	s.todos[id] = todo
	return todo
}

func (s *TodoStore) Get(id string) (*Todo, bool) {
	s.RLock()
	defer s.RUnlock()

	todo, exists := s.todos[id]
	return todo, exists
}

func (s *TodoStore) GetAll() []*Todo {
	s.RLock()
	defer s.RUnlock()

	todos := make([]*Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}
	return todos
}

func (s *TodoStore) Update(id string, completed bool) (*Todo, bool) {
	s.Lock()
	defer s.Unlock()

	if todo, exists := s.todos[id]; exists {
		todo.Completed = completed
		return todo, true
	}
	return nil, false
}

func (s *TodoStore) Delete(id string) bool {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.todos[id]; exists {
		delete(s.todos, id)
		return true
	}
	return false
}
