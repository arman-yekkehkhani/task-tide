package user

import . "github.com/arman-yekkehkhani/task-tide/internal/model/base"

type User struct {
	BaseEntity
	Username string `json:"username"`
	Password string `json:"password"`
}
