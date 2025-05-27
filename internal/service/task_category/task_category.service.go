package task_category_service

import (
	"context"

	task_category_repository "github.com/098765432m/monthly_planner_backend/internal/repository/task_category"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

type TaskCategoryService struct {
	db *task_category_repository.Queries
}

func NewTaskCategoryService(db *task_category_repository.Queries) *TaskCategoryService {
	return &TaskCategoryService{
		db: db,
	}
}

type CreateTaskCategoryServiceParams struct {
	Name        string      `json:"name"`
	Description pgtype.Text `json:"description,omitempty"`
	MonthID     pgtype.UUID `json:"month_id"`
}

func (s *TaskCategoryService) CreateTaskCategory(ctx context.Context, args *CreateTaskCategoryServiceParams) (pgtype.UUID, error) {
	var description pgtype.Text
	description.Scan(args.Description)

	id, err := s.db.CreateTaskCategory(ctx, task_category_repository.CreateTaskCategoryParams{
		Name:        args.Name,
		Description: args.Description,
		MonthID:     args.MonthID,
	})

	if err != nil {
		zap.S().Errorln(err)
		return pgtype.UUID{}, err
	}

	return id, nil
}

func (s *TaskCategoryService) DeleteTaskCategory(ctx context.Context, id pgtype.UUID) error {

	err := s.db.DeleteTaskCategory(ctx, id)
	if err != nil {
		zap.S().Errorln(err)
		return err
	}

	return nil
}
