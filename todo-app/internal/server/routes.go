package server

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"todo-app/cmd/web"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/js/*", echo.WrapHandler(fileServer))

	// e.GET("/", echo.WrapHandler(templ.Handler(web.HelloForm())))
	e.GET("/", s.RootHandler)
	e.POST("/todo", s.AddTodoHandler)
	e.DELETE("/todo/:id", s.DeleteTodoHandler)
	e.POST("/todo/:id/toggleDone", s.TooggleDoneHandler)
	// e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (s *Server) RootHandler(c echo.Context) error {
	todos := s.todo_service.GetTodos()
	return Render(c, http.StatusOK, web.ListTodos(todos))
}

func (s *Server) AddTodoHandler(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	note := strings.TrimSpace(c.Request().FormValue("note"))

	if note == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	newId := s.todo_service.AddTodo(&note)

	component := web.TodoComponent(&newId, &note)
	return Render(c, http.StatusOK, component)
}

func (s *Server) DeleteTodoHandler(c echo.Context) error {
	unsafeTodoId := c.Param("id")
	todoId, err := parseStringIdToInt(unsafeTodoId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	s.todo_service.RemoveTodo(uint(todoId))
	return Render(c, http.StatusOK, templ.NopComponent)
}

// Returns an error if the string is empty ("") or cannot be parsed
func parseStringIdToInt(requestedEventId string) (id int, err error) {
	if len(requestedEventId) == 0 {
		return -1, errors.New("ID was not supplied in request's query params")
	}
	eventId, err := strconv.Atoi(requestedEventId)

	if err != nil {
		return -1, errors.New("ID param is not a number!")
	}
	return eventId, nil
}

func (s *Server) TooggleDoneHandler(c echo.Context) error {
	unsafeTodoId := c.Param("id")
	todoId, err := parseStringIdToInt(unsafeTodoId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	todo := s.todo_service.ToggleTodoDone(uint(todoId))
	if todo == nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	id := uint(todoId)
	var component templ.Component
	if todo.Done {
		component = web.DoneTodo(&id, todo.Note)
	} else {
		component = web.TodoComponent(&id, todo.Note)
	}
	return Render(c, http.StatusOK, component)

}
