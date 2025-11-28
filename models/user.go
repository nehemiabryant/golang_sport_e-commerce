package models

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	RoleID    int
	CreatedAt time.Time
}
