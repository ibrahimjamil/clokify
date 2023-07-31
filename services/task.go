package services

import (
	. "clokify/models"
	. "clokify/types"
	"errors"
	"fmt"
	"log"
)

type TaskServiceManager struct {
	*ServiceManager
}

func (ts *TaskServiceManager) CreateTask(task *Task, srv *ServiceManager) (error, bool) {
	taskService := &TaskServiceManager{
		ServiceManager: srv,
	}

	userService := &UserServiceManager{
		ServiceManager: srv,
	}

	projectService := &ProjectServiceManager{
		ServiceManager: srv,
	}

	err, user := userService.GetUser(1)
	if err != nil {
		fmt.Print(user)
		log.Fatal("user not found for given id to create task")
	}

	err, project := projectService.GetProject(1)
	if err != nil {
		log.Fatal("user not found for given id to create task")
	}

	task.User = user
	task.Project = project

	taskErr, _ := taskService.GetTask(1)
	if taskErr == nil {
		return errors.New("task already exists"), false
	}

	res := ts.Db.Model(&Task{}).Create(&task)
	if res.Error != nil {
		return errors.New("task not create successfully"), false
	}
	return nil, true
}

func (ts *TaskServiceManager) GetTask(id int) (error, *Task) {
	var task *Task
	if err := ts.Db.Model(&task).Where("id = ?", id).First(&task); err.Error != nil {
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
