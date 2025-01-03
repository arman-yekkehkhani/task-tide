package chore

import (
	. "github.com/arman-yekkehkhani/task-tide/internal/model/base"
	model "github.com/arman-yekkehkhani/task-tide/internal/model/chore"
	"github.com/arman-yekkehkhani/task-tide/internal/repo/chore/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenTitleAndDescription_CreateChore_IsSuccessful(t *testing.T) {
	// given
	c := &model.Chore{
		Title: "title", Description: "desc",
	}

	repo := mocks.NewMockRepository(t)
	repo.EXPECT().Create(c).Return(ID(1), nil)

	svc := ServiceImpl{
		Repo: repo,
	}

	// when
	_, err := svc.Create(c)

	// then
	assert.Nil(t, err)
}

func TestGivenChore_CreateChore_shouldSetPendingStatus(t *testing.T) {
	c := &model.Chore{
		Title: "title", Description: "desc",
	}

	repo := mocks.NewMockRepository(t)
	repo.EXPECT().Create(
		&model.Chore{
			Title: "title", Description: "desc", Status: model.PENDING,
		},
	).Return(ID(1), nil)

	svc := ServiceImpl{
		Repo: repo,
	}

	// when
	_, err := svc.Create(c)

	// then
	assert.Nil(t, err)
}

func TestGivenEmptyTitle_CreateChore_ReturnsError(t *testing.T) {
	// given
	testcases := []string{"", " "}
	repo := mocks.NewMockRepository(t)
	svc := ServiceImpl{Repo: repo}

	for _, title := range testcases {
		c := &model.Chore{
			Title: title,
		}

		// when
		_, err := svc.Create(c)

		// then
		assert.EqualError(t, err, EmptyTitleOrDescription.Error())
	}
}

func TestGivenTitleAndDescription_UpdateChore_ReturnsUpdatedChore(t *testing.T) {
	// given
	title := "new title"
	description := "new description"
	id := ID(1)
	newChore := &model.Chore{
		ID:          id,
		Title:       title,
		Description: description,
	}

	oldChore := &model.Chore{
		ID:          id,
		Title:       "title",
		Description: "description",
	}

	repo := mocks.NewMockRepository(t)

	repo.EXPECT().GetByID(id).Return(oldChore)
	repo.EXPECT().Save(newChore).Return(newChore, nil)

	svc := ServiceImpl{Repo: repo}

	// when
	res, err := svc.Update(newChore)

	// then
	assert.Nil(t, err)
	assert.Equal(t, title, res.Title)
	assert.Equal(t, description, res.Description)
}

func TestGivenEmptyTitle_UpdateChore_ShouldNotChangeTile(t *testing.T) {
	// given
	oldTitle := "title"
	oldDescription := "description"
	newTitle := "  "
	newDescription := "new description"
	id := ID(1)
	newChore := &model.Chore{
		ID:          id,
		Title:       newTitle,
		Description: newDescription,
	}
	oldChore := &model.Chore{
		ID:          id,
		Title:       oldTitle,
		Description: oldDescription,
	}
	updateChore := &model.Chore{
		ID:          id,
		Title:       oldTitle,
		Description: newDescription,
	}

	repo := mocks.NewMockRepository(t)

	repo.EXPECT().GetByID(id).Return(oldChore)
	repo.EXPECT().Save(updateChore).Return(updateChore, nil)

	svc := ServiceImpl{Repo: repo}

	// when
	res, err := svc.Update(newChore)

	// then
	assert.Nil(t, err)
	assert.Equal(t, oldTitle, res.Title)
	assert.Equal(t, newDescription, res.Description)
}

func TestGivenNonExistingId_UpdateChore_ShouldReturnNotFoundError(t *testing.T) {
	// given
	id := ID(1)
	c := &model.Chore{ID: id}

	repo := mocks.NewMockRepository(t)
	repo.EXPECT().GetByID(id).Return(nil)

	svc := ServiceImpl{Repo: repo}

	// when
	_, err := svc.Update(c)

	// then
	assert.EqualError(t, err, NotFound.Error())
}

func TestGivenNewStatus_UpdateChore_ShouldSetNewStatus(t *testing.T) {
	// given
	id := ID(1)
	title := "title"
	description := "description"

	oldChore := &model.Chore{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      model.PENDING,
	}

	newChore := &model.Chore{
		ID:          id,
		Description: description,
		Status:      model.IN_PROGRESS,
	}

	updatedChore := &model.Chore{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      model.IN_PROGRESS,
	}

	repo := mocks.NewMockRepository(t)
	repo.EXPECT().GetByID(id).Return(oldChore)
	repo.EXPECT().Save(updatedChore).Return(updatedChore, nil)

	svc := ServiceImpl{Repo: repo}

	// when
	res, err := svc.Update(newChore)

	// then
	assert.Nil(t, err)
	assert.Equal(t, model.IN_PROGRESS, res.Status)
	assert.Equal(t, title, res.Title)
	assert.Equal(t, description, res.Description)
	assert.Equal(t, id, res.ID)
}

func TestGivenId_DeleteChore_ShouldCallRepoDelete(t *testing.T) {
	// given
	id := ID(1)
	c := &model.Chore{
		ID:          id,
		Title:       "title",
		Description: "desc",
	}

	repo := mocks.NewMockRepository(t)
	repo.EXPECT().DeleteById(ID(1))

	svc := ServiceImpl{
		Repo: repo,
	}

	// when
	svc.Delete(c)

	// then
}
