package chore

import (
	"errors"
	model "github.com/arman-yekkehkhani/task-tide/internal/model/chore"
	repo "github.com/arman-yekkehkhani/task-tide/internal/repo/chore"
	"strings"
)

const (
	NilChoreMsg      = "can not update using nil chore"
	NonExistingChore = "can not update using non-existent chore"
)

var (
	EmptyTitleOrDescription = errors.New("empty chore title, description")
)

type Service interface {
	Create(chore model.Chore) (model.ID, error)
}

type ServiceImpl struct {
	Repo repo.Repository
}

func (s *ServiceImpl) Create(chore *model.Chore) (model.ID, error) {
	if strings.TrimSpace(chore.Title) == "" {
		return 0, EmptyTitleOrDescription
	}

	id, err := s.Repo.Create(chore)
	return id, err
}
