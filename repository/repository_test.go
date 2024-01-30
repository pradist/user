package repository_test

import (
	"errors"
	"testing"

	"github.com/pradist/user/domain"
	"github.com/pradist/user/repository"
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
	if err != nil {
		t.Error("Expect error to be nil but got:", err)
	}
}

func TestAddUserInMemoryWithRepositoryError(t *testing.T) {
	r := NewAddDomainFakeRepository()
	r.MockAdd = func(user domain.User) error {
		return errors.New("ErrUserAlreadyExists")
	}

	sut := repository.NewAddUserInMemory(r)
	err := sut.Save(domain.User{})
	if err.Error() != errors.New("ErrUserAlreadyExists").Error() {
		t.Error("Expect error to be ErrUserAlreadyExists but got:", err)
	}
}
