package user

import (
	"io"
	"net/http"
)

func Get() error {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		return err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	println(string(b))

	return nil
}
