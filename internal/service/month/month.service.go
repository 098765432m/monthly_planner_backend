package month_service

import (
	"context"
	"fmt"
	"time"

	day_repository "github.com/098765432m/monthly_planner_backend/internal/repository/day"
	month_repository "github.com/098765432m/monthly_planner_backend/internal/repository/month"
	"github.com/098765432m/monthly_planner_backend/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

type MonthService struct {
	repo     *month_repository.Queries
	day_repo *day_repository.Queries
}

func NewMonthService(repo *month_repository.Queries, day_repo *day_repository.Queries) *MonthService {
	return &MonthService{
		repo:     repo,
		day_repo: day_repo,
	}
}

func (s *MonthService) CreateMonth(ctx context.Context, month int8, year int16) error {
	if month < 0 || month > 12 || year < 0 {
		return fmt.Errorf("month or year are not valid")
	}

	//Create new month
	monthId, err := s.repo.CreateMonth(ctx, month_repository.CreateMonthParams{
		Month: int32(month),
		Year:  int32(year),
	})

	zap.S().Infof("MonthId: %v\n", monthId)
	if err != nil {
		zap.S().Error(err)
		return err
	}

	zap.S().Infof("Month param: %#v\n", month)
	zap.S().Infof("year param: %#v\n", year)
	zap.S().Infof("Month: %#v\n", time.Month(month))
	zap.S().Infof("year: %#v\n", int(year))

	// Compute DateStart and DateEnd of a Month
	dateStart, dateEnd, err := utils.GetDateStartAndEndOfMonth(time.Month(month), int(year))
	if err != nil {
		zap.S().Error(err)
		return err
	}

	zap.S().Infof("date start: %#v\n", dateStart)
	zap.S().Infof("date end: %#v\n", dateEnd)

	//Create a range of days of above month
	err = s.day_repo.CreateRangeOfDays(ctx, day_repository.CreateRangeOfDaysParams{
		MonthID:   monthId,
		DateStart: pgtype.Date{Time: dateStart, Valid: true},
		DateEnd:   pgtype.Date{Time: dateEnd, Valid: true},
	})

	if err != nil {
		zap.S().Error(err)
		return err
	}

	return nil
}

func (s *MonthService) DeleteMonth(ctx context.Context, id pgtype.UUID) error {
	if !id.Valid {
		return fmt.Errorf("id is not valid")
	}

	return s.repo.DeleleMonth(ctx, id)
}
