package chore

import (
	"errors"
	model "github.com/arman-yekkehkhani/task-tide/internal/model/chore"
	repo "github.com/arman-yekkehkhani/task-tide/internal/repo/chore"
	"strings"
)

var (
	EmptyTitleOrDescription = errors.New("empty chore title, description")
	NotFound                = errors.New("chore does not exist")
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
		return nil, NotFound
	}

	if strings.TrimSpace(new.Title) != "" {
		old.Title = new.Title
	}
	old.Description = new.Description

	return s.Repo.Save(old)
}

func (s *ServiceImpl) Delete(chore *model.Chore) {
	s.Repo.DeleteById(chore.ID)
}
