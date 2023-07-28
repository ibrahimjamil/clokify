package models

import (
	"time"
)

type Task struct {
	ID          int
	Description string    `db:"description"`
	IsBillable  bool      `db:"is_billable"`
	tags        string    `db:"tags"`
	StartAt     time.Time `db:"start_at"`
	EndAt       time.Time `db:"end_at"`
	UserId      string    `gorm:"column:user_id;not null"`
	ProjectId   int       `gorm:"column:project_id;not null"`
	User        User      `gorm:"foreignKey:UserId"`
	Project     Project   `gorm:"foreignKey:ProjectId"`
}

func GetTaskStruct() *Task {
	return &Task{}
}
