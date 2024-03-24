package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Shteyd/wallet-app/src/backend/internal/core/contract"
	"github.com/Shteyd/wallet-app/src/backend/internal/core/dto"
	"github.com/Shteyd/wallet-app/src/backend/internal/core/entity"
	"github.com/georgysavva/scany/v2/sqlscan"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

var _ contract.UserRepository = (*Repository)(nil)

type Repository struct {
	db contract.DBTX
}

func New(db contract.DBTX) Repository {
	return Repository{
		db: db,
	}
}

const createUserQuery = "INSERT INTO users (email, secret) VALUES ($1, $2);"

func (r *Repository) CreateUser(ctx context.Context, dto dto.CreateUserDto) (entity.UserID, error) {
	const op = "repository.user.create"

	var userID entity.UserID

	err := r.db.
		QueryRowContext(ctx, createUserQuery, dto.Email, dto.Secret).
		Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return userID, nil
}

const deleteUserQuery = "DELETE FROM users WHERE id = $1;"

func (r *Repository) DeleteUser(ctx context.Context, userID entity.UserID) error {
	const op = "repository.user.delete"

	info, err := r.db.ExecContext(ctx, deleteUserQuery, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	count, err := info.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if count == 0 {
		return fmt.Errorf("%s: %w", op, ErrUserNotFound)
	}

	return nil
}

// UpdateUser implements contract.UserRepository.
func (r *Repository) UpdateUser(ctx context.Context, dto dto.UpdateUserDto) error {
	panic("unimplemented")
}

const selectUserQuery = `
SELECT
	id,
	username,
	email,
	secret,
	is_blocked,
	is_confirmed,
	is_deleted,
	created_at,
	updated_at
FROM users WHERE id = $1;
`

func (r *Repository) User(ctx context.Context, userID entity.UserID) (entity.User, error) {
	const op = "repository.user.user"

	var user userModel

	err := sqlscan.Select(ctx, r.db, &user, selectUserQuery, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}

		return entity.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user.ToEntity(), nil
}
