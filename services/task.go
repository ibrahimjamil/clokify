package services

import (
	. "clokify/models"
	. "clokify/types"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type TaskServiceManager struct {
	*ServiceManager
}

func (ts *TaskServiceManager) CreateTask(task *TaskCreateType, srv *ServiceManager) (error, bool) {
	taskService := &TaskServiceManager{
		ServiceManager: srv,
	}

	userService := &UserServiceManager{
		ServiceManager: srv,
	}

	projectService := &ProjectServiceManager{
		ServiceManager: srv,
	}

	fmt.Print(task)
	userID, err := strconv.Atoi(task.UserId)
	if err != nil {
		panic(err)
	}

	err, user := userService.GetUser(userID)
	if err != nil {
		fmt.Print(user)
		log.Fatal("user not found for given id to create task")
	}

	err, project := projectService.GetProject(task.ProjectId)
	if err != nil {
		log.Fatal("user not found for given id to create task")
	}
	DemoTask := &Task{
		ID:          task.ID,
		Description: task.Description,
		IsBillable:  task.IsBillable,
		UserId:      task.UserId,
		ProjectId:   project.ID,
	}

	taskErr, _ := taskService.GetTask(DemoTask.ID)
	if taskErr == nil {
		return errors.New("task already exists"), false
	}

	res := ts.Db.Model(&Task{}).Create(&DemoTask)
	if res.Error != nil {
		return errors.New("task not create successfully"), false
	}
	return nil, true
}

func (ts *TaskServiceManager) GetTask(id int) (error, *Task) {
	var task *Task
	if err := ts.Db.Model(&task).Preload("Project").Preload("User").Where("id = ?", id).First(&task); err.Error != nil {
		return errors.New("didnt find task"), &Task{}
	}

	return nil, task
}

func (ts *TaskServiceManager) DeleteTask(id int) (error, Task) {
	var task Task
	if err := ts.Db.Model(&task).Where("id = ?", id).Delete(&task); err.Error != nil {
		return errors.New("delete failed"), Task{}
	}

	return nil, task
}
