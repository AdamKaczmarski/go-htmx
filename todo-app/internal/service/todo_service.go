package service

import (
	"fmt"
	"sync"
)

type Service interface {
	GetTodos() map[uint]todo
	AddTodo(note *string) uint
	MarkTodoAsDone(id uint)
	RemoveTodo(id uint)
}

type todo struct {
	done bool
	note *string
}

type service struct {
	mu      *sync.Mutex
	todos   map[uint]todo
	last_id *uint
}

func NewService() *service {
	todos := make(map[uint]todo)
	var last_id uint = 1
	return &service{&sync.Mutex{}, todos, &last_id}
}

func (s *service) GetTodos() map[uint]todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	clonedMap := make(map[uint]todo, len(s.todos))
	for k, v := range s.todos {
		clonedMap[k] = v
	}
	return clonedMap
}

func (s *service) AddTodo(note *string) uint {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := *s.last_id
	*s.last_id += 1
	newTodo := todo{
		done: false,
		note: note,
	}
	s.todos[id] = newTodo
	return id
}
func (s *service) MarkTodoAsDone(id uint) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if todo, ok := s.todos[id]; ok {
		todo.done = true
		s.todos[id] = todo
	}
}

func (s *service) RemoveTodo(id uint) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.todos, id)
}

func (t todo) String() string {
	return fmt.Sprintf("Todo{note: \"%s\", done: %v}", *t.note, t.done)
}
