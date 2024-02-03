package repository_test

import (
	"testing"

	"github.com/pradist/user/domain"
	"github.com/pradist/user/repository"
	"github.com/stretchr/testify/assert"
)

type UserInMemoryMockRepository struct {
	MockGetFn  func() ([]domain.User, error)
	MockSaveFn func(domain.User) error
}

func (fake *UserInMemoryMockRepository) Get() ([]domain.User, error) {
	return fake.MockGetFn()
}

func (fake *UserInMemoryMockRepository) Save(user domain.User) error {
	return fake.MockSaveFn(user)
}

func newUserInMemoryMockRepository() *UserInMemoryMockRepository {
	return &UserInMemoryMockRepository{
		MockGetFn:  func() ([]domain.User, error) { return nil, nil },
		MockSaveFn: func(user domain.User) error { return nil },
	}
}

func TestGetUserInMemorySucceed(t *testing.T) {
	r := newUserInMemoryMockRepository()
	sut := repository.NewUserInMemoryRepository(r)

	_, err := sut.Get()
	assert.Nil(t, err)
}

func TestSaveUserInMemorySucceed(t *testing.T) {
	r := newUserInMemoryMockRepository()
	sut := repository.NewUserInMemoryRepository(r)

	err := sut.Save(domain.User{})
	assert.Nil(t, err)
}
