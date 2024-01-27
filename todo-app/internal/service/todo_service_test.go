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
        if expected != int(*s.getServiceData().last_id) {
			panic(fmt.Sprintf("expected: %d, actual: %d", expected, *s.getServiceData().last_id))
		}
	})
	t.Run("AddingTodos", func(t *testing.T) {
		s := NewService()
		expected := make(map[uint]Todo)
		for i := 1; i < 6; i++ {
			expected[uint(i)] = Todo{false, &note}
			_ = s.AddTodo(&note)
		}
		if !reflect.DeepEqual(expected, s.getServiceData().todos) {
			panic(fmt.Sprintf("notes, not equal, expected: %+v, actual %+v", expected, s.getServiceData().todos))
		}
	})
	t.Run("MarkingTodoAsDone", func(t *testing.T) {
		s := NewService()
		str := "doneNote"
		id := s.AddTodo(&str)
		s.ToggleTodoDone(id)
		for _, todo := range s.getServiceData().todos {
			if !todo.Done {
				panic(fmt.Sprintf("Todo should be done, actual %v", todo))
			}
		}
		s.ToggleTodoDone(id)
		for _, todo := range s.getServiceData().todos {
			if todo.Done {
				panic(fmt.Sprintf("Todo should be NOT done, actual %v", todo))
			}
		}
	})
	t.Run("RemovingTodo", func(t *testing.T) {
		s := NewService()
		str := "toBeRemoved"
		id := s.AddTodo(&str)
		if len(s.getServiceData().todos) != 1 {
			panic("Todo was not added")
		}
		s.RemoveTodo(id)
		if len(s.getServiceData().todos) != 0 {
			panic(fmt.Sprintf("Todo was not removed %v", s.getServiceData().todos))
		}
	})
}
