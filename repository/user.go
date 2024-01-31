package repository

import "github.com/pradist/user/domain"

type InMemoryRepository interface {
	Get() ([]domain.User, error)
}

type InMemoryUser struct {
	repository InMemoryRepository
}

func (usecase *InMemoryUser) Get() ([]domain.User, error) {
	users, err := usecase.repository.Get()
	return users, err
}

func NewInMemoryUser(r InMemoryRepository) *InMemoryUser {
	return &InMemoryUser{repository: r}
}
