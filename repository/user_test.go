package repository_test

import (
	"testing"

	"github.com/pradist/user/domain"
	"github.com/pradist/user/repository"
	"github.com/stretchr/testify/assert"
)

type UserInMemoryMockRepository struct {
	MockGetFn    func() ([]domain.User, error)
	MockSaveFn   func(domain.User) error
	MockUpdateFn func(domain.User) error
	MockDeleteFn func(int) error
}

func (fake *UserInMemoryMockRepository) Get() ([]domain.User, error) {
	return fake.MockGetFn()
}

func (fake *UserInMemoryMockRepository) Save(user domain.User) error {
	return fake.MockSaveFn(user)
}

func (fake *UserInMemoryMockRepository) Update(user domain.User) error {
	return fake.MockUpdateFn(user)
}

func (fake *UserInMemoryMockRepository) Delete(id int) error {
	return fake.MockDeleteFn(id)
}

func newUserInMemoryMockRepository() *UserInMemoryMockRepository {
	return &UserInMemoryMockRepository{
		MockGetFn:    func() ([]domain.User, error) { return nil, nil },
		MockSaveFn:   func(user domain.User) error { return nil },
		MockUpdateFn: func(user domain.User) error { return nil },
		MockDeleteFn: func(id int) error { return nil },
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

func TestUpdateUserInMemorySucceed(t *testing.T) {
	r := newUserInMemoryMockRepository()
	sut := repository.NewUserInMemoryRepository(r)

	err := sut.Update(domain.User{})

	assert.Nil(t, err)
}

func TestUpdateUserInMemoryFail(t *testing.T) {
	r := newUserInMemoryMockRepository()
	sut := repository.NewUserInMemoryRepository(r)

	err := sut.Update(domain.User{ID: 99})

	assert.NotNil(t, err)
}

func TestDeleteUserInMemorySucceed(t *testing.T) {
	r := newUserInMemoryMockRepository()
	sut := repository.NewUserInMemoryRepository(r)

	err := sut.Delete(1)

	assert.Nil(t, err)
}

func TestDeleteUserInMemoryFail(t *testing.T) {
	r := newUserInMemoryMockRepository()
	sut := repository.NewUserInMemoryRepository(r)

	err := sut.Delete(99)

	assert.NotNil(t, err)
}
