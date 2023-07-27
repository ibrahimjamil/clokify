package db

import (
	. "clokify/models"
	"fmt"

	"gorm.io/gorm"
)

func MigrateAllDB(db *gorm.DB) {
	userMigrationErr := db.AutoMigrate(&User{})
	if userMigrationErr != nil {
		fmt.Printf("User didnt able to migrate")
	}

	projectMigrationErr := db.AutoMigrate(&Project{})
	if projectMigrationErr != nil {
		fmt.Printf("project didnt able to migrate")
	}

	taskMigrationErr := db.AutoMigrate(&Task{})
	if taskMigrationErr != nil {
		fmt.Printf("Task didnt able to migrate")
	}
}
