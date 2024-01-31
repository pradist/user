package usecase_test

import (
	"testing"

	"github.com/pradist/user/domain"
	"github.com/pradist/user/usecase"
)

type addEventFakeRepository struct {
	MockAddFn func(domain.Event) error
}

func (fake *addEventFakeRepository) Add(e domain.Event) error {
	return fake.MockAddFn(e)
}

func newAddEventFakeRepository() *addEventFakeRepository {
	return &addEventFakeRepository{
		MockAddFn: func(e domain.Event) error { return nil },
	}
}

func TestAddEventInMemorySucceed(t *testing.T) {
	r := newAddEventFakeRepository()
	sut := usecase.NewAddEventInMemory(r)

	err := sut.Save(domain.Event{})
	if err != nil {
		t.Error("Expect error to be nil but got:", err)
	}
}
