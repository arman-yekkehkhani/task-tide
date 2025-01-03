package user

import (
	"errors"
	. "github.com/arman-yekkehkhani/task-tide/internal/model/base"
	. "github.com/arman-yekkehkhani/task-tide/internal/model/user"
	"github.com/arman-yekkehkhani/task-tide/internal/repo/user/mocks"
	"github.com/arman-yekkehkhani/task-tide/internal/service/security"
	securityMock "github.com/arman-yekkehkhani/task-tide/internal/service/security/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenUsernameAndPassword_CreateUser_ReturnsUserWithHashedPassword(t *testing.T) {
	// given
	username := "user"
	password := "pass"
	hashedPass := "hashed"

	hashSvc := securityMock.NewMockHashService(t)
	hashSvc.EXPECT().Hash(security.BCRYPT, password).Return(hashedPass, nil)

	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().GetByUsername(username).Return(nil, UsernameNotFound)
	repo.EXPECT().Create(
		&User{
			Username: username,
			Password: hashedPass,
		}).Return(
		&User{
			Username: username,
			Password: hashedPass,
			BaseEntity: BaseEntity{
				ID:        ID(1),
				IsDeleted: false,
			},
		}, nil)

	svc := ServiceImpl{
		repo:        repo,
		hashService: hashSvc,
	}

	// when
	user, err := svc.Create(username, password)

	// then
	assert.Nil(t, err)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, hashedPass, user.Password)
	assert.Equal(t, false, user.IsDeleted)
}

func TestGivenExistingUsername_CreateUser_ReturnsError(t *testing.T) {
	// given
	username := "user"
	password := "pass"

	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().GetByUsername(username).Return(&User{}, nil)

	svc := ServiceImpl{
		repo: repo,
	}
	// when
	user, err := svc.Create(username, password)

	// then
	assert.EqualError(t, err, UsernameAlreadyExists.Error())
	assert.Nil(t, user)
}

func TestGivenEmptyOrWhiteSpacePassword_CreateUser_ReturnsError(t *testing.T) {
	// given
	passwords := []string{
		"", "  ",
	}

	svc := ServiceImpl{
		repo:        nil,
		hashService: nil,
	}

	// when
	for _, password := range passwords {
		_, err := svc.Create("username", password)
		// then
		assert.NotNil(t, err, EmptyPassword)
	}
}

func TestGivenHashServiceError_CreateUser_ReturnsErrorAndAbortUserCreation(t *testing.T) {
	// given
	username := "user"
	password := "pass"

	hashSvc := securityMock.NewMockHashService(t)
	hashingErr := errors.New("random error")
	hashSvc.EXPECT().Hash(security.BCRYPT, password).Return("", hashingErr)

	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().GetByUsername(username).Return(nil, UsernameNotFound)

	svc := ServiceImpl{
		repo:        repo,
		hashService: hashSvc,
	}

	// when
	_, err := svc.Create(username, password)

	// then
	assert.EqualError(t, err, hashingErr.Error())
}