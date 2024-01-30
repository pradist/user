package repository_test

import (
	"errors"
	"testing"

	"github.com/pradist/user/domain"
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

func TestAddMemory(t *testing.T) {
	repo := NewAddDomainFakeRepository()
	repo.MockAdd = func(user domain.User) error {
		return nil
	}

	err := repo.Add(domain.User{})
	assert.NoError(t, err)
}

func TestAddMemoryWithErr(t *testing.T) {
	repo := NewAddDomainFakeRepository()
	repo.MockAdd = func(user domain.User) error {
		return errors.New("some error")
	}

	err := repo.Add(domain.User{})
	assert.Error(t, err)
	assert.EqualError(t, err, "some error")
}
