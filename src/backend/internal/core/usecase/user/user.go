package user

import (
	"context"
	"fmt"

	"github.com/Shteyd/wallet-app/src/backend/internal/core/contract"
	"github.com/Shteyd/wallet-app/src/backend/internal/core/dto"
	"github.com/Shteyd/wallet-app/src/backend/internal/core/entity"
)

type Usecase struct {
	repoManager contract.RepositoryManager
	hasher      contract.SecretHasher
}

func NewUsecase(
	repoManager contract.RepositoryManager,
	hasher contract.SecretHasher,
) Usecase {
	return Usecase{
		repoManager: repoManager,
		hasher:      hasher,
	}
}

type SaveUserDto struct {
	UserID   entity.UserID
	Username string
	Email    string
	Secret   string
}

func (uc Usecase) Save(ctx context.Context, info SaveUserDto) (entity.User, error) {
	const op = "usecase.user.save"

	secretHash, err := uc.hasher.Hash(info.Secret)
	if err != nil {
		return entity.User{}, fmt.Errorf("%s: %w", op, err)
	}

	repo := uc.repoManager.UserRepo()

	if info.UserID.IsEmpty() {
		info.UserID, err = repo.CreateUser(ctx, dto.NewCreateUserDto(info.Email, secretHash))
	} else {
		err = repo.UpdateUser(ctx, dto.NewUpdateUserDto(info.UserID, info.Username, info.Email, secretHash))
	}

	if err != nil {
		return entity.User{}, fmt.Errorf("%s: %w", op, err)
	}

	user, err := repo.User(ctx, info.UserID)
	if err != nil {
		return entity.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
