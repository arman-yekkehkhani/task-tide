package main

import (
	"fmt"
	"github.com/arman-yekkehkhani/task-tide/server"
)

func main() {
	s := server.New()
	s.Register()
	s.Start()

	fmt.Println("tasktide started")
}
