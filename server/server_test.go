package server

import (
	"net/http"
	"testing"
)

type mockHandler struct{}

func (h mockHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func TestRegister_AddsChoreCreate(t *testing.T) {
	//server := New(mockHandler{})
	//server.Register()
	//server.Start()
	//
	//buffer := bytes.Buffer{}
	//http.NewRequest(http.MethodPost, "/chores", &buffer)

}
