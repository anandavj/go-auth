package service

import (
	"context"
	"errors"
	"fmt"
	"go-auth/internal/helper/log"
	"go-auth/internal/model"
	"go-auth/internal/repository"
	"reflect"
	"runtime"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Empty struct{}

var pckg = reflect.TypeOf(Empty{}).PkgPath() //Get Package Name

type UserServiceImpl struct {
	userRepository repository.UserRepository
	logger         *zap.Logger
}

func NewUserServiceImpl(userRepository repository.UserRepository, logger *zap.Logger) UserService {
	return &UserServiceImpl{userRepository, logger}
}

func (service *UserServiceImpl) CreateNewUser(ctx context.Context, user *model.User) error {
	// Logging Purpose
	pc, _, _, _ := runtime.Caller(0) //Get Function Name
	method := runtime.FuncForPC(pc).Name()
	service.logger.Info("### Creating New User...", log.Field(pckg, method, fmt.Sprintf("ID: %v Email: %v Username %v", user.ID, user.Email, user.Username), nil)...)

	// Validation Required Field
	validate := validator.New()

	if validateErr := validate.Struct(user); validateErr != nil {
		service.logger.Error("### Failed Creating New User - Validation", log.Field(pckg, method, "", validateErr)...)
		return validateErr
	}

	// Check if User Exist
	isUserExist, err := service.isUserExist(ctx, user)
	if isUserExist {
		userExistErr := errors.New("user already exist")
		service.logger.Error("### Failed Creating New User - User Already Exist", log.Field(pckg, method, "", userExistErr)...)
		return userExistErr
	}

	if err != nil {
		service.logger.Error("### Failed Creating New User", log.Field(pckg, method, "", err)...)
		return err
	}

	createErr := service.userRepository.CreateNewUser(ctx, user)

	if createErr != nil {
		service.logger.Error("### Failed Creating New User", log.Field(pckg, method, "", createErr)...)
		return err
	}

	service.logger.Info("### Success Creating New User", log.Field(pckg, method, "", nil)...)
	return nil
}

func (service *UserServiceImpl) isUserExist(ctx context.Context, user *model.User) (isExist bool, err error) {
	// Logging Purpose
	pc, _, _, _ := runtime.Caller(0) //Get Function Name
	method := runtime.FuncForPC(pc).Name()
	service.logger.Info("### Checking if User Exist...", log.Field(pckg, method, fmt.Sprintf("ID: %v Email: %v Username %v", user.ID, user.Email, user.Username), nil)...)

	isIdExistRes, isIdExistErr := service.userRepository.IsUserIdExist(ctx, user.ID)
	isEmailExistRes, isEmailExistErr := service.userRepository.IsUserEmailExist(ctx, user.Email)
	isUsernameExistRes, isUsernameExistErr := service.userRepository.IsUserUsernameExist(ctx, user.Username)

	if !isIdExistRes {
		service.logger.Error("### Failed Checking if User ID Exist", log.Field(pckg, method, "", isIdExistErr)...)
		if isIdExistErr != nil {
			err = isIdExistErr
		}
		return isExist, err
	}
	if !isEmailExistRes {
		service.logger.Error("### Failed Checking if User Email Exist", log.Field(pckg, method, "", isEmailExistErr)...)
		if isEmailExistErr != nil {
			err = isEmailExistErr
		}
		return isExist, err
	}
	if !isUsernameExistRes {
		service.logger.Error("### Failed Checking if User Username Exist", log.Field(pckg, method, "", isUsernameExistErr)...)
		if isUsernameExistErr != nil {
			err = isUsernameExistErr
		}
		return isExist, err
	}
	service.logger.Info("### Success Checking if User Exist", log.Field(pckg, method, fmt.Sprintf("ID: %v Email: %v Username %v", user.ID, user.Email, user.Username), nil)...)
	return true, nil
}
