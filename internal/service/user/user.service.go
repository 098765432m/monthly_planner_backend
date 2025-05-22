package user_service

import (
	"context"
	"fmt"

	user_repository "github.com/098765432m/monthly_planner_backend/internal/repository/user"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *user_repository.Queries
}

func NewUserService(repo *user_repository.Queries) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUserById(ctx context.Context, idParam string) (user_repository.User, error) {
	var id pgtype.UUID
	err := id.Scan(idParam)
	if err != nil {
		zap.S().Errorln("Invalid User UUID!")
		return user_repository.User{}, err
	}

	return s.repo.GetUserById(ctx, id)
}

type CreateUserServiceParams struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (s *UserService) CreateUser(ctx context.Context, params *CreateUserServiceParams) (user_repository.User, error) {

	if params.Username == "" || params.Password == "" || params.Email == "" || params.PhoneNumber == "" {
		zap.S().Errorln("Missing out params")
		return user_repository.User{}, fmt.Errorf("field cannot empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.S().Errorln("Failed to hashed password")
		return user_repository.User{}, err
	}

	params.Password = string(hashedPassword)

	return s.repo.CreateUser(ctx, user_repository.CreateUserParams{
		Username:    params.Username,
		Password:    params.Password,
		Email:       params.Email,
		PhoneNumber: params.PhoneNumber,
	})
}

type UpdateUserByIdServiceParams struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	IsActive    bool   `json:"is_active"`
}

func (s *UserService) UpdateUserById(ctx context.Context, idParam string, args *UpdateUserByIdServiceParams) error {
	var id pgtype.UUID
	err := id.Scan(idParam)
	if err != nil {

		zap.S().Errorln("Invalid User UUID!")
		return err
	}

	if args.Username == "" || args.Password == "" || args.Email == "" || args.PhoneNumber == "" {
		zap.S().Errorln("Missing out params!")
		return fmt.Errorf("field cannot empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(args.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.S().Errorln("Failed to hashed password!")
		return err
	}

	args.Password = string(hashedPassword)
	var role user_repository.RoleEnum
	if err := role.Scan(args.Role); err != nil {
		zap.S().Errorln("Unsupported role enum!")
		return err
	}

	err = s.repo.UpdateUser(ctx, user_repository.UpdateUserParams{
		UserID:      id,
		Username:    args.Username,
		Password:    args.Password,
		Email:       args.Email,
		PhoneNumber: args.PhoneNumber,
		Role:        role,
		IsActive:    args.IsActive,
	})

	if err != nil {
		zap.S().Errorln(err)
		return err
	}

	return nil
}

func (s *UserService) DeleteUserById(ctx context.Context, idParam string) error {
	var id pgtype.UUID
	err := id.Scan(idParam)
	if err != nil {
		zap.S().Errorln("Invalid User UUID!")
		return err
	}

	return s.repo.DeleteUser(ctx, id)
}
