package entity

import "time"

type User struct {
	ID          int64
	Username    string
	Email       string
	Secret      string
	IsBlocked   bool
	IsConfirmed bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
