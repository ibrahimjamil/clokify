package services

import (
	. "clokify/models"
	. "clokify/types"
	"errors"
	"log"
)

type ProjectServiceManager struct {
	*ServiceManager
}

func (ps *ProjectServiceManager) CreateProject(project *ProjectCreateType, userId int, srv *ServiceManager) (error, bool) {
	userService := &UserServiceManager{
		ServiceManager: srv,
	}

	err, user := userService.GetUser(userId)
	if err != nil {
		return errors.New("user not found for given id to create project"), false
	}

	DemoProject := &Project{
		ID:       project.ID,
		Name:     project.Name,
		IsPublic: project.IsPublic,
		ColorTag: project.ColorTag,
	}
	DemoProject.Users = append(DemoProject.Users, user)

	projectErr, _ := ps.GetProjectByName(DemoProject.Name)
	if projectErr == nil {
		return errors.New("project already exists"), false
	}

	res := ps.Db.Model(&Project{}).Create(&DemoProject)
	if res.Error != nil {
		return errors.New("project not create successfully"), false
	}

	return nil, true
}

func (ps *ProjectServiceManager) GetProjectByName(name string) (error, Project) {
	var project Project
	if err := ps.Db.Model(&project).Where("name = ?", name).First(&project); err.Error != nil {
		return errors.New("didnt find project"), Project{}
	}

	return nil, project
}

func (ps *ProjectServiceManager) GetProject(id int) (error, Project) {
	var project Project
	if err := ps.Db.Model(&project).Preload("Users").Where("id = ?", id).First(&project); err.Error != nil {
		return errors.New("didnt find project"), Project{}
	}

	return nil, project
}

func (ps *ProjectServiceManager) DeleteProject(id int) (error, Project) {
	var project Project

	err, fetchedProject := ps.GetProject(id)
	if err != nil {
		return errors.New("didnt get project"), Project{}
	}

	// here transaction is used for the purpose of chaining some queries that
	// we dont want to execute if one of the query in between begin and commit fails.
	tx := ps.Db.Begin()

	if err := tx.First(&project, &Project{ID: fetchedProject.ID}).Error; err != nil {
		tx.Rollback()
		log.Fatal("failed to get the project")
	}

	if err := tx.Exec("DELETE FROM project_users WHERE project_id = ?", project.ID).Error; err != nil {
		tx.Rollback()
		log.Fatal("failed to delete junction table rows")
	}

	if err := tx.Delete(&project).Error; err != nil {
		tx.Rollback()
		log.Fatal("failed to delete the project")
	}

	tx.Commit()

	return nil, project
}
