package chore

import (
	"errors"
)

const (
	NilChoreMsg      = "can not update using nil chore"
	NonExistingChore = "can not update using non-existent chore"
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

func (s *ServiceImpl) Update(chore *Chore) (*Chore, error) {
	if chore == nil {
		return nil, errors.New(NilChoreMsg)
	}

	c := s.Repo.GetByID(chore.ID)
	if c == nil {
		return nil, errors.New(NonExistingChore)
	}
	//
	//if chore.Title != "" {
	//	c.Title = chore.Title
	//}
	//
	//if chore.Description != "" {
	//	c.Description = chore.Description
	//}
	//
	c, err := s.Repo.Update(c)
	return c, err
}
