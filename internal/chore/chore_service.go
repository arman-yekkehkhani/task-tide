package chore

import (
	"errors"
)

var (
	EmptyTitleOrDescription = errors.New("empty chore title, description")
)

type Service interface {
	Create(chore Chore) (ID, error)
}

type ServiceImpl struct {
	Repo Repository
}

func (s *ServiceImpl) Create(chore Chore) (ID, error) {
	if len(chore.Title) == 0 || len(chore.Description) == 0 {
		return -1, EmptyTitleOrDescription
	}

	id, err := s.Repo.Create(&chore)
	if err != nil {
		return 0, err
	}

	return id, nil
}
