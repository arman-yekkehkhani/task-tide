package chore

import (
	"encoding/json"
	"fmt"
	chore2 "github.com/arman-yekkehkhani/task-tide/internal/model/chore"
	chore3 "github.com/arman-yekkehkhani/task-tide/internal/service/chore"
	"net/http"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type HandlerImpl struct {
	Srv chore3.Service
}

func (h HandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	c := chore2.Chore{}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.Srv.Create(c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%d", id)))
}
