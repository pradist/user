package repository

import (
	"errors"

	"github.com/pradist/user/domain"
)

var users = []domain.User{
	{ID: 1, Name: "John"},
	{ID: 2, Name: "Doe"},
}

type UserInMemoryRepository interface {
	Get() ([]domain.User, error)
	Save(domain.User) error
	Update(domain.User) error
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

func (usecase *UserInMemory) Update(u domain.User) error {
	for i, user := range users {
		if user.ID == u.ID {
			users[i] = u
			return nil
		}
	}
	return errors.New("user not found")
}

func (usecase *UserInMemory) Delete(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func NewUserInMemoryRepository(r UserInMemoryRepository) *UserInMemory {
	return &UserInMemory{repository: r}
}
