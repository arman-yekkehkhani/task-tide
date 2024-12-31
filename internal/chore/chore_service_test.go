package chore

import (
	"errors"
	"testing"
)

type MockRepo struct {
	storage map[ID]Chore
}

func (r MockRepo) Create(chore Chore) error {
	r.storage[chore.ID] = chore
	return nil
}

func (MockRepo) Update(chore Chore) error {
	//TODO implement me
	panic("implement me")
}

func (MockRepo) Delete(chore Chore) error {
	//TODO implement me
	panic("implement me")
}

func (MockRepo) GetById(id int32) (*Chore, error) {
	//TODO implement me
	panic("implement me")
}

func TestCreateChore_Successful(t *testing.T) {
	mockRepo := MockRepo{storage: make(map[ID]Chore)}

	chore := Chore{
		Title:       "title",
		Description: "description",
	}

	service := ServiceImpl{
		Repo: mockRepo,
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
