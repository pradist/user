package repository

import (
	"github.com/pradist/user/domain"
)

type AddInMemoryRepository interface {
	Add(domain.User) error
}

type addInMemoryRepository struct{}

func NewAddInMemoryRepository() AddInMemoryRepository {
	return &addInMemoryRepository{}
}

func (repo *addInMemoryRepository) Add(user domain.User) error {
	return nil
}
