package cruds

import (
	. "clokify/models"
	. "clokify/services"
	"log"

	. "clokify/types"

	"gorm.io/gorm"
)

func TaskCrud(db *gorm.DB, srvMan *ServiceManager) {
	taskService := &TaskServiceManager{
		ServiceManager: srvMan,
	}

	// create task
	task := &TaskCreateType{
		ID:          1,
		Description: "Enxsys task",
		IsBillable:  true,
	}
	err, taskRes := taskService.CreateTask(task, srvMan)

	if err == nil {
		log.Println("task created succesfully", taskRes)
	} else {
		log.Println(err)
	}

	// get task
	err, getProject := taskService.GetTask(task.ID)
	if err == nil {
		log.Println("task fetched successfully", getProject)
	} else {
		log.Println(err, getProject)
	}

	// delete task
	// err, deletedProject := taskService.DeleteTask(task.ID)
	// if err == nil {
	// 	log.Println("task deleted successfully", deletedProject)
	// } else {
	// 	log.Println(err, deletedProject)
	// }
}
