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

func (ps *ProjectServiceManager) CreateProject(project *Project, srv *ServiceManager) (error, bool) {
	userService := &UserServiceManager{
		ServiceManager: srv,
	}

	err, user := userService.GetUser(1)
	if err != nil {
		log.Fatal("user not found for given id to create project")
	}
	project.Users = append(project.Users, user)

	projectErr, _ := ps.GetProjectByName(project.Name)
	if projectErr == nil {
		return errors.New("project already exists"), false
	}

	res := ps.Db.Model(&Project{}).Create(&project)
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
	if err := ps.Db.Model(&project).Where("id = ?", id).First(&project); err.Error != nil {
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
