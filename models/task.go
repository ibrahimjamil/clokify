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

type TaskCreateType struct {
	ID          int    `json:"ID" binding:"required"`
	Description string `json:"Description" binding:"required"`
	IsBillable  bool   `json:"IsBillable" binding:"required"`
	UserId      string `json:"UserId" binding:"required"`
	ProjectId   int    `json:"ProjectId" binding:"required"`
}
