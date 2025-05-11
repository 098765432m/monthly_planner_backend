package task_service

import (
	"context"
	"fmt"

	task_repository "github.com/098765432m/monthly_planner_backend/internal/repository/task"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

type TaskService struct {
	repo *task_repository.Queries
}

func NewTaskService(repo *task_repository.Queries) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, arg task_repository.CreateTaskParams) (task_repository.Task, error) {

	return s.repo.CreateTask(ctx, arg)
}

func (s *TaskService) DeleteTaskById(ctx context.Context, id pgtype.UUID) error {
	if !id.Valid {
		zap.S().Errorln("id is not valid")
		return fmt.Errorf("id is not valid")
	}

	return s.repo.DeleteTaskById(ctx, id)
}

func (s *TaskService) UpdateTaskById(ctx context.Context, arg task_repository.UpdateTaskByIdParams) (task_repository.Task, error) {
	if !arg.ID.Valid {
		zap.S().Errorln("id is not valid")
		return task_repository.Task{}, fmt.Errorf("id is not valid")
	}

	return s.repo.UpdateTaskById(ctx, arg)
}

func (s *TaskService) GetAllTasksOfADay(ctx context.Context) {

}
