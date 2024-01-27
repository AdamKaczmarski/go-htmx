package service

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTodoService(t *testing.T) {
	note := "note"
	t.Run("IncrementingLastId", func(t *testing.T) {
		s := NewService()
		expected := 1
		for i := 0; i < 6; i++ {
			_ = s.AddTodo(&note)
			expected += 1
		}
		if expected != int(*s.last_id) {
			panic(fmt.Sprintf("expected: %d, actual: %d", expected, *s.last_id))
		}
	})
	t.Run("AddingTodos", func(t *testing.T) {
		s := NewService()
		expected := make(map[uint]todo)
		for i := 1; i < 6; i++ {
			expected[uint(i)] = todo{false, &note}
			_ = s.AddTodo(&note)
		}
		if !reflect.DeepEqual(expected, s.todos) {
			panic(fmt.Sprintf("notes, not equal, expected: %+v, actual %+v", expected, s.todos))
		}
	})
	t.Run("MarkingTodoAsDone", func(t *testing.T) {
		s := NewService()
		str := "doneNote"
		id := s.AddTodo(&str)
		s.MarkTodoAsDone(id)
		for _, todo := range s.todos {
			if !todo.done {
				panic(fmt.Sprintf("Todo should be done, actual %v", todo))
			}
		}
	})
	t.Run("RemovingTodo", func(t *testing.T) {
		s := NewService()
		str := "toBeRemoved"
		id := s.AddTodo(&str)
		if len(s.todos) != 1 {
			panic("Todo was not added")
		}
		s.RemoveTodo(id)
		if len(s.todos) != 0 {
			panic(fmt.Sprintf("Todo was not removed %v", s.todos))
		}
	})
}
