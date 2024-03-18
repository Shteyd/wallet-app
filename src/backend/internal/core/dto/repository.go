package dto

import "github.com/Shteyd/wallet-app/src/backend/internal/core/entity"

type CreateUserDto struct {
	Email  string
	Secret string
}

func NewCreateUserDto(email, secret string) CreateUserDto {
	return CreateUserDto{
		Email:  email,
		Secret: secret,
	}
}

type UpdateUserDto struct {
	ID       entity.UserID
	Username string
	Email    string
	Secret   string
}

func NewUpdateUserDto(
	userID entity.UserID,
	username, email, secret string,
) UpdateUserDto {
	return UpdateUserDto{
		ID:       userID,
		Username: username,
		Email:    email,
		Secret:   secret,
	}
}
