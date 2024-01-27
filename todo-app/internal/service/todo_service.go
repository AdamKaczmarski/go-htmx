package service

import (
	"fmt"
	"sync"
)

type Service interface {
	GetTodos() map[uint]Todo
	AddTodo(note *string) uint
	ToggleTodoDone(id uint) *Todo
	RemoveTodo(id uint)
    GetTodo(id uint) *Todo
    getServiceData() *service
}

type Todo struct {
	Done bool
	Note *string
}

type service struct {
	mu      *sync.Mutex
	todos   map[uint]Todo
	last_id *uint
}

func NewService() Service {
	todos := make(map[uint]Todo)
	var last_id uint = 1
	return &service{&sync.Mutex{}, todos, &last_id}
}

func (s *service) getServiceData() *service{
    return s
}
func (s *service) GetTodos() map[uint]Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	clonedMap := make(map[uint]Todo, len(s.todos))
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
	newTodo := Todo{
		Done: false,
		Note: note,
	}
	s.todos[id] = newTodo
	return id
}
func (s *service) ToggleTodoDone(id uint) *Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	if todo, ok := s.todos[id]; ok {
		todo.Done = !todo.Done 
		s.todos[id] = todo
        return &todo 
	}
    return nil
}

func (s *service) RemoveTodo(id uint) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.todos, id)
}

func (s *service) GetTodo(id uint) *Todo{
	s.mu.Lock()
	defer s.mu.Unlock()
    todo := s.todos[id]
	return &todo 
}

func (t Todo) String() string {
	return fmt.Sprintf("Todo{note: \"%s\", done: %v}", *t.Note, t.Done)
}
