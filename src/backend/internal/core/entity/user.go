package entity

import (
	"time"
)

type UserID int64

func (id UserID) IsEmpty() bool {
	return !(id > 0)
}

type User struct {
	ID          UserID
	Username    string
	Email       string
	Secret      string
	IsBlocked   bool
	IsConfirmed bool
	IsDeleted   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
