package repository

import (
	"context"
	"go-auth/internal/model"
)

type UserRepository interface {
	CreateNewUser(context.Context, *model.User) error
	IsUserIdExist(context.Context, string) (bool, error)
	IsUserEmailExist(context.Context, string) (bool, error)
	IsUserUsernameExist(context.Context, string) (bool, error)
}
