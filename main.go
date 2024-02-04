package main

import (
	"fmt"

	"github.com/pradist/user/domain"
	"github.com/pradist/user/repository"
)

func main() {

	r := repository.NewUserInMemoryRepository(nil)
	users, _ := r.Get()

	fmt.Printf("%v\n", users)

	u := domain.User{ID: 3, Name: "Jane"}

	r.Save(u)

	users, _ = r.Get()

	fmt.Printf("%v\n", users)
}
