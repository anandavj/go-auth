package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-auth/internal/helper"
	"go-auth/internal/helper/log"
	"go-auth/internal/model"
	"reflect"
	"runtime"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UserRepositoryImpl struct {
	db     *sqlx.DB
	logger *zap.Logger
}

const TABLE_NAME = "user"

type Empty struct{}

var pckg = reflect.TypeOf(Empty{}).PkgPath() //Get Package Name

func NewUserRepositoryImpl(conn *sqlx.DB, logger *zap.Logger) UserRepository {
	return &UserRepositoryImpl{conn, logger}
}

func (repo *UserRepositoryImpl) CreateNewUser(ctx context.Context, user *model.User) error {
	// Logging Purpose
	pc, _, _, _ := runtime.Caller(0) //Get Function Name
	method := runtime.FuncForPC(pc).Name()
	repo.logger.Info("### Creating User...", log.Field(pckg, method, fmt.Sprintf("Email: %v", user.Email), nil)...)

	query := `
		INSERT INTO $1 (
			id,
			email,
			username,
			password,
			refresh_token,
			first_name,
			last_name,
			created_at,
			created_by,
			modified_at,
			modified_by,
			password_fail,
			last_success_login,
			last_fail_login,
			last_changed_password,
			security_code
		)
		VALUES ($2,$3,$4,$5,NULL,$6,$7,$8,$9,$8,$10,0,$8,$8,$8,NULL)
	`
	currentTime := helper.GetCurrentTime()
	result, err := repo.db.ExecContext(ctx, query, TABLE_NAME, user.ID, user.Email, user.Username, user.Password, user.FirstName, user.LastName, currentTime, user.CreatedBy, user.ModifiedBy)

	if err != nil {
		repo.logger.Error("### Failed Creating User", log.Field(pckg, method, "", err)...)
		return err
	}
	repo.logger.Info("### Success Creating User", log.Field(pckg, method, fmt.Sprintf("ID: %v", result), nil)...)
	return nil
}

func (repo *UserRepositoryImpl) IsUserIdExist(ctx context.Context, id string) (bool, error) {
	// Logging Purpose
	pc, _, _, _ := runtime.Caller(0) //Get Function Name
	method := runtime.FuncForPC(pc).Name()
	repo.logger.Info("### Checking if User Exist By ID...", log.Field(pckg, method, fmt.Sprintf("ID: %v ", id), nil)...)

	query := `
		SELECT id
		FROM $1
		WHERE id = $2
	`

	user := &model.User{}

	err := repo.db.QueryRowContext(ctx, query, TABLE_NAME, id).Scan(&user.ID)
	if err != nil {
		repo.logger.Error("### Failed Checking User By Id", log.Field(pckg, method, "", err)...)
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	repo.logger.Info("### Success Checking User By Id", log.Field(pckg, method, fmt.Sprintf("ID: %v ", id), nil)...)
	return true, nil
}

func (repo *UserRepositoryImpl) IsUserEmailExist(ctx context.Context, email string) (bool, error) {
	// Logging Purpose
	pc, _, _, _ := runtime.Caller(0) //Get Function Name
	method := runtime.FuncForPC(pc).Name()
	repo.logger.Info("### Checking if User Exist By Email...", log.Field(pckg, method, fmt.Sprintf("Email: %v ", email), nil)...)

	query := `
		SELECT id
		FROM $1
		WHERE email = $2
	`

	user := &model.User{}

	err := repo.db.QueryRowContext(ctx, query, TABLE_NAME, email).Scan(&user.ID)
	if err != nil {
		repo.logger.Error("### Failed Checking User By Email", log.Field(pckg, method, "", err)...)
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	repo.logger.Info("### Success Checking User By Email", log.Field(pckg, method, fmt.Sprintf("Email: %v ", email), nil)...)
	return true, nil
}

func (repo *UserRepositoryImpl) IsUserUsernameExist(ctx context.Context, username string) (bool, error) {
	// Logging Purpose
	pc, _, _, _ := runtime.Caller(0) //Get Function Name
	method := runtime.FuncForPC(pc).Name()
	repo.logger.Info("### Checking if User Exist By Username...", log.Field(pckg, method, fmt.Sprintf("Username: %v ", username), nil)...)

	query := `
		SELECT id
		FROM $1
		WHERE username = $2
	`

	user := &model.User{}

	err := repo.db.QueryRowContext(ctx, query, TABLE_NAME, username).Scan(&user.ID)
	if err != nil {
		repo.logger.Error("### Failed Checking User By Username", log.Field(pckg, method, "", err)...)
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	repo.logger.Info("### Success Checking User By Username", log.Field(pckg, method, fmt.Sprintf("Username: %v ", username), nil)...)
	return true, nil
}
