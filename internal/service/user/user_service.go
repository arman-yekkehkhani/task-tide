package user

import (
	"errors"
	. "github.com/arman-yekkehkhani/task-tide/internal/model/user"
	. "github.com/arman-yekkehkhani/task-tide/internal/repo/user"
	"github.com/arman-yekkehkhani/task-tide/internal/service/security"
	"strings"
)

var (
	UsernameAlreadyExists = errors.New("username already exists")
	UsernameNotFound      = errors.New("username does not exist")
	EmptyPassword         = errors.New("password is empty or whitespace")
)

type ServiceImpl struct {
	repo        UserRepository
	hashService security.HashService
}

func (s *ServiceImpl) Create(username string, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, EmptyPassword
	}

	if _, err := s.repo.GetByUsername(username); err == nil {
		return nil, UsernameAlreadyExists
	}

	hashedPass, err := s.hashService.Hash(security.BCRYPT, password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: username,
		Password: hashedPass,
	}
	saved, err := s.repo.Create(user)
	return saved, err
}
