//go:build integration

package user_test

import (
	"testing"

	"github.com/pradist/user/user"
)

func TestIntegrationGetUser(t *testing.T) {
	err := user.Get()

	if err != nil {
		t.Error(err)
	}
}
