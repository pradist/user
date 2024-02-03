package repository

import "github.com/pradist/user/domain"

var users = []domain.User{
	{ID: 1, Name: "John"},
	{ID: 2, Name: "Doe"},
}

type UserInMemoryRepository interface {
	Get() ([]domain.User, error)
	Save(domain.User) error
}

type UserInMemory struct {
	repository UserInMemoryRepository
}

func (usecase *UserInMemory) Get() ([]domain.User, error) {
	return users, nil
}

func (usecase *UserInMemory) Save(u domain.User) error {
	users = append(users, u)
	return nil
}

func NewUserInMemoryRepository(r UserInMemoryRepository) *UserInMemory {
	return &UserInMemory{repository: r}
}
