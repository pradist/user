package repository

import (
	"errors"

	"github.com/pradist/user/domain"
)

var ErrUserAlreadyExists = errors.New("ErrUserAlreadyExists")

type AddInMemoryRepository interface {
	Add(domain.User) error
}

type AddUserInMemory struct {
	repository AddInMemoryRepository
}

func (a *AddUserInMemory) Save(user domain.User) error {
	return a.repository.Add(user)
}

func NewAddUserInMemory(repository AddInMemoryRepository) *AddUserInMemory {
	return &AddUserInMemory{repository: repository}
}
