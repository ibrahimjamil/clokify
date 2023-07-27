package models

import (
	"time"
)

type Task struct {
	ID            int
	Description   string        `db:"description"`
	IsBillable    bool          `db:"is_billable"`
	tags          string        `db:"tags"`
	StartTime     time.Time     `db:"start_time"`
	EndTime       time.Time     `db:"end_time"`
	EstimatedTime time.Duration `db:"estimated_time"`
	UserId        string        `db:"user_id"`
	ProjectId     string        `db:"project_id"`
	User          User          `gorm:"foreignKey:user_id"`
	Project       Project       `gorm:"foreignKey:project_id"`
}

func GetTaskStruct() *Task {
	return &Task{}
}
