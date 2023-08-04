package services

import (
	. "clokify/models"
	. "clokify/types"
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type TaskServiceManager struct {
	*ServiceManager
}

type TaskResult struct {
	err  error
	task Task
}

func (ts *TaskServiceManager) CreateTask(task *TaskCreateType, srv *ServiceManager) (error, bool) {

	userService := &UserServiceManager{
		ServiceManager: srv,
	}

	projectService := &ProjectServiceManager{
		ServiceManager: srv,
	}

	taskService := &TaskServiceManager{
		ServiceManager: srv,
	}

	userID, err := strconv.Atoi(task.UserId)
	if err != nil {
		return errors.New("userId issue"), false
	}

	userCh := make(chan UserResult)
	projectCh := make(chan ProjectResult)
	taskCh := make(chan TaskResult)

	var wg sync.WaitGroup
	fmt.Print("before")

	wg.Add(1)
	go userService.GetUserWithChannel(userID, userCh, &wg)

	wg.Add(1)
	go projectService.GetProjectWithChannel(task.ProjectId, projectCh, &wg)

	wg.Add(1)
	go taskService.GetTaskWithChannel(task.ID, taskCh, &wg)

	userResult := <-userCh
	projectResult := <-projectCh
	taskResult := <-taskCh

	wg.Wait()
	close(userCh)
	close(projectCh)
	close(taskCh)

	fmt.Print("after")

	if userResult.err != nil {
		fmt.Print(userResult)
		return errors.New("user not found for given id to create task"), false
	}

	if projectResult.err != nil {
		return errors.New("project not found for given id to create task"), false
	}

	if taskResult.err == nil {
		fmt.Print(taskResult)
		return errors.New("task already exists"), false
	}

	DemoTask := &Task{
		ID:          task.ID,
		Description: task.Description,
		IsBillable:  task.IsBillable,
		UserId:      task.UserId,
		ProjectId:   projectResult.project.ID,
	}

	res := ts.Db.Model(&Task{}).Create(&DemoTask)
	if res.Error != nil {
		return errors.New("task not create successfully"), false
	}
	return nil, true
}

func (ts *TaskServiceManager) GetTask(id int) (error, Task) {
	var task Task
	if err := ts.Db.Model(&task).Preload("Project").Preload("User").Where("id = ?", id).First(&task); err.Error != nil {
		return errors.New("didnt find task"), Task{}
	}

	return nil, task
}

func (ts *TaskServiceManager) GetTaskWithChannel(
	id int,
	ch chan TaskResult,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	taskResult := TaskResult{}
	taskResult.err, taskResult.task = ts.GetTask(id)
	ch <- taskResult
	return
}

func (ts *TaskServiceManager) DeleteTask(id int) (error, Task) {
	var task Task
	if err := ts.Db.Model(&task).Where("id = ?", id).Delete(&task); err.Error != nil {
		return errors.New("delete failed"), Task{}
	}

	return nil, task
}
