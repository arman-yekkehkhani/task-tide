package chore

import (
	"errors"
	"testing"
)

type InMemoryRepository struct {
	db []Chore
}

func (r *InMemoryRepository) Create(chore *Chore) (ID, error) {
	if r.db == nil {
		r.db = make([]Chore, 0)
	}
	chore.ID = ID(len(r.db) + 1)
	r.db = append(r.db, *chore)
	return chore.ID, nil
}

func (r *InMemoryRepository) Update(chore *Chore) error {
	//TODO implement me
	panic("implement me")
}

func (r *InMemoryRepository) Delete(chore *Chore) error {
	//TODO implement me
	panic("implement me")
}

func (r *InMemoryRepository) GetById(id int32) (*Chore, error) {
	//TODO implement me
	panic("implement me")
}

func TestCreateChore_Successful(t *testing.T) {
	repo := InMemoryRepository{}

	chore := Chore{
		Title:       "title",
		Description: "description",
	}

	service := ServiceImpl{
		Repo: &repo,
	}

	_, err := service.Create(chore)

	if err != nil {
		t.Errorf("expected %s, got %s", "nil", err)
	}
}

func TestCreateChore_WhenEmptyTitle_ShouldReturnErr(t *testing.T) {
	chore := Chore{
		Title:       "",
		Description: "description",
	}

	service := ServiceImpl{}

	_, err := service.Create(chore)

	if err == nil || !errors.Is(err, EmptyTitleOrDescription) {
		t.Errorf("expected %s, got %s", EmptyTitleOrDescription, err)
	}
}
