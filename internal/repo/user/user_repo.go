package user

import (
	. "github.com/arman-yekkehkhani/task-tide/internal/model/user"
	. "github.com/arman-yekkehkhani/task-tide/internal/repo/crud"
)

type UserRepository interface {
	CrudRepository[User]
	GetByUsername(username string) (*User, error)
}
