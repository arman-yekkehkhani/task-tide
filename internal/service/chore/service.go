package chore

import (
	"errors"
	model "github.com/arman-yekkehkhani/task-tide/internal/model/chore"
	repo "github.com/arman-yekkehkhani/task-tide/internal/repo/chore"
	"strings"
)

const ()

var (
	EmptyTitleOrDescription = errors.New("empty chore title, description")
	ChoreNotFound           = errors.New("chore does not exist")
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

func (s *ServiceImpl) Update(new *model.Chore) (*model.Chore, error) {
	old := s.Repo.GetByID(new.ID)
	if old == nil {
		return nil, ChoreNotFound
	}

	if strings.TrimSpace(new.Title) != "" {
		old.Title = new.Title
	}
	old.Description = new.Description

	return s.Repo.Save(old)
}
