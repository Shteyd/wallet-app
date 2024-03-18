package contract

import (
	"context"

	"github.com/Shteyd/wallet-app/src/backend/internal/core/dto"
	"github.com/Shteyd/wallet-app/src/backend/internal/core/entity"
)

type (
	RepositoryManager interface {
		Begin(ctx context.Context) (RepositoryManager, error)
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error

		UserRepo() UserRepository
	}

	UserRepository interface {
		CreateUser(ctx context.Context, dto dto.CreateUserDto) (entity.UserID, error)
		UpdateUser(ctx context.Context, dto dto.UpdateUserDto) error
		DeleteUser(ctx context.Context, userID entity.UserID) error

		User(ctx context.Context, userID entity.UserID) (entity.User, error)
	}
)
