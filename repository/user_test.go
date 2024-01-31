package repository_test

import (
	"testing"

	"github.com/pradist/user/domain"
	"github.com/pradist/user/repository"
)

type InMemoryMockRepository struct {
	MockGetFn func() ([]domain.User, error)
}

func (fake *InMemoryMockRepository) Get() ([]domain.User, error) {
	return fake.MockGetFn()
}

func newInMemoryMockRepository() *InMemoryMockRepository {
	return &InMemoryMockRepository{
		MockGetFn: func() ([]domain.User, error) { return nil, nil },
	}
}

func TestGetUserInMemorySucceed(t *testing.T) {
	r := newInMemoryMockRepository()
	sut := repository.NewInMemoryUser(r)

	_, err := sut.Get()
	if err != nil {
		t.Error("Expect error to be nil but got:", err)
	}
}
