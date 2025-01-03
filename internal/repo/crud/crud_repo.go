package crud

import "github.com/arman-yekkehkhani/task-tide/internal/model/base"

type CrudRepository[T any] interface {
	Create(t *T) (*T, error)
	Save(t *T) (*T, error)
	GetById(id base.ID) (*T, error)
	DeleteById()
}
