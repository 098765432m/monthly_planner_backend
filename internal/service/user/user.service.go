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

func (s *UserService) GetUserById(ctx context.Context, id pgtype.UUID) (user_repository.User, error) {
	if !id.Valid {
		return user_repository.User{}, fmt.Errorf("id is not valid")
	}

	return s.repo.GetUserById(ctx, id)
}

func (s *UserService) CreateUser(ctx context.Context, params user_repository.CreateUserParams) (user_repository.User, error) {
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

	return s.repo.CreateUser(ctx, params)
}

func (s *UserService) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	if !id.Valid {
		return fmt.Errorf("id is not valid")
	}

	return s.repo.DeleteUser(ctx, id)
}
