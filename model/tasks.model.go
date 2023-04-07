package model

import "time"

// ID added for mongoDB
type Task struct {
	Title       string
	Description string
	User        string
	IsDone      bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
