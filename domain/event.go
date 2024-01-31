package domain

import "time"

type Event struct {
	Title    string
	Date     time.Time
	Place    string
	Category string
	KeyWords []string
}
