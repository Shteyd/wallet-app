package user

import (
	"time"

	"github.com/Shteyd/wallet-app/src/backend/internal/core/entity"
)

type userModel struct {
	ID          int64     `db:"id"`
	Username    string    `db:"username"`
	Email       string    `db:"email"`
	Secret      string    `db:"secret"`
	IsBlocked   bool      `db:"is_blocked"`
	IsConfirmed bool      `db:"is_confirmed"`
	IsDeleted   bool      `db:"is_deleted"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (m userModel) ToEntity() entity.User {
	return entity.User{
		ID:          entity.UserID(m.ID),
		Username:    m.Username,
		Email:       m.Email,
		Secret:      m.Secret,
		IsBlocked:   m.IsBlocked,
		IsConfirmed: m.IsConfirmed,
		IsDeleted:   m.IsDeleted,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
