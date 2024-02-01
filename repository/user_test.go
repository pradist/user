package repository_test

import (
	"testing"

	"github.com/pradist/user/domain"
	"github.com/pradist/user/repository"
)

type UserInMemoryMockRepository struct {
	MockGetFn func() ([]domain.User, error)
}

func (fake *UserInMemoryMockRepository) Get() ([]domain.User, error) {
	return fake.MockGetFn()
}

func newUserInMemoryMockRepository() *UserInMemoryMockRepository {
	return &UserInMemoryMockRepository{
		MockGetFn: func() ([]domain.User, error) { return nil, nil },
	}
}

func TestGetUserInMemorySucceed(t *testing.T) {
	r := newUserInMemoryMockRepository()
	sut := repository.NewUserInMemoryRepository(r)

	_, err := sut.Get()
	if err != nil {
		t.Error("Expect error to be nil but got:", err)
	}
}
