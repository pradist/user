package main

import (
	"fmt"

	"github.com/pradist/user/domain"
	"github.com/pradist/user/repository"
)

func main() {

	r := repository.NewUserInMemoryRepository(nil)
	users, _ := r.Get()

	fmt.Printf("get all %v\n", users)

	u := domain.User{ID: 3, Name: "Jane"}

	r.Save(u)

	users, _ = r.Get()
	fmt.Printf("after save %v\n", users)

	r.Update(domain.User{ID: 3, Name: "Jane Doe"})

	users, _ = r.Get()

	fmt.Printf("after update %v\n", users)

	r.Delete(3)

	users, _ = r.Get()
	fmt.Printf("after delete %v\n", users)
}
