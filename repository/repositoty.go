package repository

import (
	"github.com/pradist/user/domain"
)

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
