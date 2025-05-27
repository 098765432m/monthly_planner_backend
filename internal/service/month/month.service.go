package month_service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	day_repository "github.com/098765432m/monthly_planner_backend/internal/repository/day"
	month_repository "github.com/098765432m/monthly_planner_backend/internal/repository/month"
	task_repository "github.com/098765432m/monthly_planner_backend/internal/repository/task"
	"github.com/098765432m/monthly_planner_backend/internal/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

type MonthService struct {
	repo      *month_repository.Queries
	task_repo *task_repository.Queries
	day_repo  *day_repository.Queries
}

type MonthServiceDesp struct {
	Repo      *month_repository.Queries
	Task_Repo *task_repository.Queries
	Day_repo  *day_repository.Queries
}

func NewMonthService(monthServiceParams *MonthServiceDesp) *MonthService {
	return &MonthService{
		repo:      monthServiceParams.Repo,
		task_repo: monthServiceParams.Task_Repo,
		day_repo:  monthServiceParams.Day_repo,
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

// SELECT t.id, t.name, t.description, t.status, t.time_start, t.time_end, d.date
type GetTaskOfMonthResponse struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description,omitempty"`
	Status         string    `json:"status"`
	TimeStart      time.Time `json:"time_start,omitempty"`
	TimeEnd        time.Time `json:"time_end,omitempty"`
	Date           time.Time `json:"day_id"`
	TaskCategoryID string    `json:"task_category_id,omitempty"`
}

// Get All task from month
func (s *MonthService) GetAllTasksOfMonth(ctx context.Context, monthId pgtype.UUID) ([]GetTaskOfMonthResponse, error) {

	month, err := s.repo.GetMonthById(ctx, monthId)
	if err != nil {
		zap.S().Error(err)
		return nil, fmt.Errorf("failed to get month by Id")
	}

	tasks, err := s.repo.GetAllTasksOfMonth(ctx, month_repository.GetAllTasksOfMonthParams{
		MonthNumber: month.Month,
		YearNumber:  month.Year,
	})
	if err != nil {
		zap.S().Error(err)
		return nil, fmt.Errorf("failed to get all tasks of the month")
	}

	var resp []GetTaskOfMonthResponse

	for _, task := range tasks {

		taskResp := GetTaskOfMonthResponse{}
		taskResp.ID = task.ID.String()
		taskResp.Name = task.Name
		taskResp.Description = task.Description.String
		taskResp.Status = string(task.Status)
		taskResp.TimeStart = time.UnixMicro(task.TimeStart.Microseconds)
		taskResp.TimeEnd = time.UnixMicro(task.TimeEnd.Microseconds)
		taskResp.Date = task.Date.Time

		resp = append(resp, taskResp)
	}

	return resp, nil
}

type TaskDays struct {
	Task []task_repository.UpdateTaskByIdParams `json:"task,omitempty"`
	Date pgtype.Date                            `json:"date"`
}

// Save all tasks of days of month
func (s *MonthService) SaveAllTaskOfMonth(ctx context.Context, month int, year int, daysTasks []TaskDays) error {

	// Check Month Exists
	checkMonth, err := s.repo.GetMonthByMonthAndYear(ctx, month_repository.GetMonthByMonthAndYearParams{
		Month: int32(month),
		Year:  int32(year),
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			//Handle if checkMonth not found or null

			// create month if not exists
			newMonthId, err := s.repo.CreateMonth(ctx, month_repository.CreateMonthParams{
				Month: int32(month),
				Year:  int32(year),
			})
			if err != nil {
				return err
			}

			checkMonth.ID = newMonthId // Add new created Id month to month

		} else {
			// Handler all other error
			return err
		}
	}

	// TODO Use goroutine efficiently and Transaction
	for _, day := range daysTasks {
		// Loop Every day check exist and every task every day

		dayRow, err := s.day_repo.GetDayByDate(ctx, day.Date) // Check Days Exists
		date := day.Date.Time
		dayOfWeek := date.Weekday() // Get WeekDay

		if errors.Is(err, pgx.ErrNoRows) {
			// Create day if not exists
			creatdeDayId, err := s.day_repo.CreateDay(ctx, day_repository.CreateDayParams{
				Date:      day.Date,
				DayOfWeek: int32(dayOfWeek),
				MonthID:   checkMonth.ID,
			})

			dayRow.ID = creatdeDayId

			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		// Check Job Exists and update if not created it
		for _, task := range day.Task {
			var status task_repository.NullTaskStatusEnum
			if err = status.Scan(task.Status); err != nil {
				status.TaskStatusEnum = task_repository.TaskStatusEnumNOTDONE
				status.Valid = true
			}
			zap.S().Info(status)
			if task.TaskID.Valid {
				// Update task
				_, err = s.task_repo.UpdateTaskById(ctx, task_repository.UpdateTaskByIdParams{
					TaskID:         task.TaskID,
					Name:           task.Name,
					Description:    task.Description,
					Status:         status,
					TimeStart:      task.TimeStart,
					TimeEnd:        task.TimeEnd,
					DayID:          task.DayID,
					TaskCategoryID: task.TaskCategoryID,
				})
			} else {
				// Create task if not exist
				_, err = s.task_repo.CreateTask(ctx, task_repository.CreateTaskParams{
					Name:           task.Name,
					Description:    task.Description,
					Status:         status,
					TimeStart:      task.TimeStart,
					TimeEnd:        task.TimeEnd,
					DayID:          dayRow.ID,
					TaskCategoryID: task.TaskCategoryID,
				})
			}

			if err != nil {
				zap.S().Errorln(err)
				return err
			}
		}
	}
	return nil
}
