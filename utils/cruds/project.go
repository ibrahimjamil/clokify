package cruds

import (
	. "clokify/models"
	. "clokify/services"
	"log"

	. "clokify/types"

	"gorm.io/gorm"
)

func ProjectCrud(db *gorm.DB, srvMan *ServiceManager) {
	projectService := &ProjectServiceManager{
		ServiceManager: srvMan,
	}

	// create project
	project := &Project{
		ID:       1,
		Name:     "Enxys",
		IsPublic: true,
		ColorTag: "#fffff",
	}
	err, projectRes := projectService.CreateProject(project, srvMan)

	if err == nil {
		log.Println("project created succesfully", projectRes)
	} else {
		log.Println(err)
	}

	// get Project
	err, getProject := projectService.GetProject(project.ID)
	if err == nil {
		log.Println("project fetched successfully", getProject)
	} else {
		log.Println(err, getProject)
	}

	// delete project
	// err, deletedProject := projectService.DeleteProject(project.ID)
	// if err == nil {
	// 	log.Println("project deleted successfully", deletedProject)
	// } else {
	// 	log.Println(err, deletedProject)
	// }
}
