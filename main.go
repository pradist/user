package main

import (
	"fmt"

	"github.com/pradist/user/repository"
)

func main() {

	r := repository.NewUserInMemoryRepository(nil)
	users, _ := r.Get()

	fmt.Printf("%v\n", users)
}
