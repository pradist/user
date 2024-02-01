package repository

import "github.com/pradist/user/domain"

type UserInMemoryRepository interface {
	Get() ([]domain.User, error)
}

type UserInMemory struct {
	repository UserInMemoryRepository
}

func (usecase *UserInMemory) Get() ([]domain.User, error) {
	users, err := usecase.repository.Get()
	return users, err
}

func NewUserInMemoryRepository(r UserInMemoryRepository) *UserInMemory {
	return &UserInMemory{repository: r}
}
