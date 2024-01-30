package repository_test

import (
	"errors"
	"testing"

	"github.com/pradist/user/domain"
	"github.com/pradist/user/repository"
	"github.com/stretchr/testify/assert"
)

type addDomainFakeRepository struct {
	MockAdd func(domain.User) error
}

func (fake *addDomainFakeRepository) Add(user domain.User) error {
	return fake.MockAdd(user)
}

func NewAddDomainFakeRepository() *addDomainFakeRepository {
	return &addDomainFakeRepository{}
}

func TestAddUserInMemory(t *testing.T) {
	r := NewAddDomainFakeRepository()
	r.MockAdd = func(user domain.User) error {
		return nil
	}

	sut := repository.NewAddUserInMemory(r)
	err := sut.Save(domain.User{})
	assert.NoError(t, err)
}

func TestAddUserInMemoryWithRepositoryError(t *testing.T) {
	r := NewAddDomainFakeRepository()
	r.MockAdd = func(user domain.User) error {
		return errors.New("ErrUserAlreadyExists")
	}

	sut := repository.NewAddUserInMemory(r)
	err := sut.Save(domain.User{})

	assert.Error(t, err)
	assert.EqualError(t, err, "ErrUserAlreadyExists")
}
