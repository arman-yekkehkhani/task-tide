package chore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) Update(c *Chore) (*Chore, error) {
	args := m.Called(c)
	return args.Get(0).(*Chore), nil
}

func (m *mockRepository) GetByID(id ID) *Chore {
	args := m.Called(id)
	chore, _ := args.Get(0).(*Chore)
	return chore
}

func (m *mockRepository) Create(chore *Chore) (ID, error) {
	args := m.Called(chore)
	return args.Get(0).(ID), args.Error(1)
}

func TestCreateChore_Successful(t *testing.T) {
	// given
	chore := Chore{
		Title:       "title",
		Description: "description",
	}
	repo := &mockRepository{}
	service := ServiceImpl{
		Repo: repo,
	}
	repo.On("Create", &chore).Return(ID(1), nil)

	// when
	_, err := service.Create(chore)

	// then
	if err != nil {
		t.Errorf("expected %s, got %s", "nil", err)
	}
}

func TestCreateChore_WhenEmptyTitle_ShouldReturnErr(t *testing.T) {
	// given
	chore := Chore{
		Title:       "",
		Description: "description",
	}
	repo := &mockRepository{}
	service := ServiceImpl{Repo: repo}

	// when
	_, err := service.Create(chore)

	// then
	assert.Error(t, err, EmptyTitleOrDescription)
}

func TestUpdateChore_Successful(t *testing.T) {
	// given
	oldChore := &Chore{ID(1), "title", "description"}
	newChore := &Chore{ID(1), "new_title", "new_description"}

	m := &mockRepository{}
	m.On("GetByID", ID(1)).Return(oldChore)
	m.On("Update", mock.Anything).Return(newChore, nil)

	svc := ServiceImpl{Repo: m}

	// when
	res, err := svc.Update(newChore)

	// then
	assert.NoError(t, err)
	assert.Equal(t, newChore.Title, res.Title)
	assert.Equal(t, newChore.Description, res.Description)
	assert.Equal(t, oldChore.ID, res.ID)

	m.AssertExpectations(t)
}

func TestUpdateNilChore_ReturnsError(t *testing.T) {
	// given
	m := &mockRepository{}
	svc := ServiceImpl{}

	// when
	_, err := svc.Update(nil)

	// then
	assert.EqualError(t, err, NilChoreMsg)
	m.AssertNotCalled(t, "Update", mock.Anything)
	m.AssertExpectations(t)
}

func TestUpdateNonExistingChore_ReturnsError(t *testing.T) {
	// given
	m := &mockRepository{}
	m.On("GetByID", ID(1)).Return(nil)
	svc := ServiceImpl{m}

	// when
	_, err := svc.Update(&Chore{ID: ID(1)})

	// then
	assert.EqualError(t, err, NonExistingChore)
	m.AssertNotCalled(t, "Update", mock.Anything)
	m.AssertExpectations(t)
}
