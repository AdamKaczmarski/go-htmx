package service

type Service interface {
	GetTodos() []todo
	AddTodo(note *string) int
	MarkTodoAsDone(id int)
    RemoveTodo(id int)
}

type todo struct {
	id   int
	done bool
	note *string
}

type service struct {
	todos   *[]todo
	last_id *int
}

func NewService() *service {
	todos := make([]todo, 10)
	last_id := 1
	return &service{&todos, &last_id}
}

func (s *service) GetTodos() []todo {
	var cloneTodos []todo = make([]todo, len(*s.todos))
	copy(cloneTodos, *s.todos)
	return cloneTodos
}

func (s *service) AddTodo(note *string) int {
	id := *s.last_id
	id += 1
	newTodo := todo{
		id:   id,
		done: false,
		note: note,
	}
	tmp := append(*s.todos, newTodo)
	s.todos = &tmp

	return id
}
func (s *service) MarkTodoAsDone(id int) {
	panic("todo")
}

func (s *service) RemoveTodo(id int) {

}
