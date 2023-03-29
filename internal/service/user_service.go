package service

import (
	"context"
	"go-auth/internal/model"
)

type UserService interface {
	CreateNewUser(context.Context, *model.User) error
}
