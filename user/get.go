package user

import (
	"encoding/json"
	"io"
	"net/http"
)

func Get() ([]User, error) {
	c := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/users", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var u []User
	err = json.Unmarshal(b, &u)

	return u, err
}
