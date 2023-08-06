package db

import (
	. "clokify/models"
	"fmt"

	"gorm.io/gorm"
)

func MigrateAllDB(db *gorm.DB) {
	userMigrationErr := db.AutoMigrate(GetUserStruct())
	if userMigrationErr != nil {
		fmt.Printf("User didnt able to migrate")
	}

	projectMigrationErr := db.AutoMigrate(GetProjectStruct())
	if projectMigrationErr != nil {
		fmt.Printf("project didnt able to migrate")
	}

	taskMigrationErr := db.AutoMigrate(GetTaskStruct())
	if taskMigrationErr != nil {
		fmt.Printf("Task didnt able to migrate")
	}
}
