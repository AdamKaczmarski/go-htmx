package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"todo-app/internal/service"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port         int
	todo_service service.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	s := service.NewService()
	for i := 0; i < 5; i++ {
		data := fmt.Sprintf("Note %d", i)
		s.AddTodo(&data)
	}
    s.ToggleTodoDone(1)
	NewServer := &Server{
		port:         port,
		todo_service: s,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
