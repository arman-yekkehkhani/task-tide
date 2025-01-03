package chore

import (
	"errors"
	. "github.com/arman-yekkehkhani/task-tide/internal/model/base"
	. "github.com/arman-yekkehkhani/task-tide/internal/model/chore"
	. "github.com/arman-yekkehkhani/task-tide/internal/repo/chore"

	"strings"
)

var (
	EmptyTitleOrDescription = errors.New("empty chore title, description")
	NotFound                = errors.New("chore does not exist")
)

type Service interface {
	Create(chore Chore) (ID, error)
}

type ServiceImpl struct {
	Repo Repository
}

func (s *ServiceImpl) Create(chore *Chore) (ID, error) {
	if strings.TrimSpace(chore.Title) == "" {
		return 0, EmptyTitleOrDescription
	}
	chore.Status = PENDING

	id, err := s.Repo.Create(chore)
	return id, err
}

func (s *ServiceImpl) Update(new *Chore) (*Chore, error) {
	old := s.Repo.GetByID(new.ID)
	if old == nil {
		return nil, NotFound
	}

	if strings.TrimSpace(new.Title) != "" {
		old.Title = new.Title
	}
	old.Description = new.Description
	old.Status = new.Status

	return s.Repo.Save(old)
}

func (s *ServiceImpl) Delete(chore *Chore) {
	s.Repo.DeleteById(chore.ID)
}
