package chore

import (
	"github.com/arman-yekkehkhani/task-tide/internal/model/chore"
	"github.com/arman-yekkehkhani/task-tide/internal/repo/chore/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenTitleAndDescription_CreateChore_IsSuccessful(t *testing.T) {
	// given
	c := &chore.Chore{
		Title: "title", Description: "desc",
	}

	repo := mocks.NewRepository(t)
	repo.EXPECT().Create(c).Return(chore.ID(1), nil)

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
	repo := mocks.NewRepository(t)
	svc := ServiceImpl{Repo: repo}

	for _, title := range testcases {
		c := &chore.Chore{
			Title: title,
		}

		// when
		_, err := svc.Create(c)

		// then
		assert.EqualError(t, err, EmptyTitleOrDescription.Error())
	}
}
