package day_service

import (
	"context"

	day_repository "github.com/098765432m/monthly_planner_backend/internal/repository/day"
)

type DayService struct {
	repo *day_repository.Queries
}

func NewDayService(repo *day_repository.Queries) *DayService {
	return &DayService{
		repo: repo,
	}
}

func (s *DayService) CreateDay(ctx context.Context, arg day_repository.CreateDayParams) error {
	return s.repo.CreateDay(ctx, arg)
}

func (s *DayService) CreateRangeOfDays(ctx context.Context, arg day_repository.CreateRangeOfDaysParams) error {
	return s.repo.CreateRangeOfDays(ctx, arg)
}

func (s *DayService) UpdateDayById(ctx context.Context, arg day_repository.UpdateDayByIdParams) error {
	return s.repo.UpdateDayById(ctx, arg)
}
