package server

import "C"
import (
	"github.com/arman-yekkehkhani/task-tide/internal/chore"
	"net/http"
)

type Server struct {
	chore chore.Handler
}

func New() *Server {
	handler := chore.HandlerImpl{
		Srv: &chore.ServiceImpl{
			Repo: &chore.SqliteRepository{},
		},
	}
	return &Server{chore: handler}
}

func (s *Server) Register() {
	http.HandleFunc("/chores", s.chore.Create)
}

func (s *Server) Start() {
	http.ListenAndServe(":8080", nil)
}
