// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package month_repository

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type TaskStatusEnum string

const (
	TaskStatusEnumNOTDONE TaskStatusEnum = "NOT_DONE"
	TaskStatusEnumDONE    TaskStatusEnum = "DONE"
)

func (e *TaskStatusEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TaskStatusEnum(s)
	case string:
		*e = TaskStatusEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for TaskStatusEnum: %T", src)
	}
	return nil
}

type NullTaskStatusEnum struct {
	TaskStatusEnum TaskStatusEnum `json:"task_status_enum"`
	Valid          bool           `json:"valid"` // Valid is true if TaskStatusEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTaskStatusEnum) Scan(value interface{}) error {
	if value == nil {
		ns.TaskStatusEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TaskStatusEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTaskStatusEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TaskStatusEnum), nil
}

type Day struct {
	ID        pgtype.UUID `json:"id"`
	Date      pgtype.Date `json:"date"`
	DayOfWeek int32       `json:"day_of_week"`
	MonthID   pgtype.UUID `json:"month_id"`
}

type Month struct {
	ID        pgtype.UUID        `json:"id"`
	Year      int32              `json:"year"`
	Month     int32              `json:"month"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type Task struct {
	ID             pgtype.UUID    `json:"id"`
	Name           string         `json:"name"`
	Description    pgtype.Text    `json:"description"`
	Status         TaskStatusEnum `json:"status"`
	TimeStart      pgtype.Time    `json:"time_start"`
	TimeEnd        pgtype.Time    `json:"time_end"`
	DayID          pgtype.UUID    `json:"day_id"`
	TaskCategoryID pgtype.UUID    `json:"task_category_id"`
}
