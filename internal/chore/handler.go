package chore

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type HandlerImpl struct {
	Srv Service
}

func (h HandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	c := Chore{}

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
