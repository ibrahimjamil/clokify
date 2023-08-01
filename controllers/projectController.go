package controllers

import (
	. "clokify/models"
	. "clokify/services"
	"encoding/json"
	"io"
	"log"
	"strconv"

	. "clokify/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProject(srvMananger *ServiceManager, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var project ProjectCreateType
		projectService := &ProjectServiceManager{
			ServiceManager: srvMananger,
		}

		bodyData, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Failed to read request body",
			})
			return
		}

		if err := json.Unmarshal(bodyData, &project); err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		err, projectRes := projectService.CreateProject(&project, project.UserId, srvMananger)
		if err == nil {
			log.Println("project created succesfully", projectRes)
			ctx.JSON(200, gin.H{"success": "project created succesfully"})
		} else {
			ctx.JSON(400, gin.H{
				"error":  "project didnt created succesfully",
				"reason": err.Error(),
			})
		}
	}
}

func GetProject(srvMananger *ServiceManager, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		projectId := ctx.Param("projectId")
		id, err := strconv.Atoi(projectId)
		if err != nil {
			// ... handle error
			panic(err)
		}

		projectService := &ProjectServiceManager{
			ServiceManager: srvMananger,
		}

		err, getProject := projectService.GetProject(id)
		if err == nil {
			log.Println("project fetched successfully", getProject)
			ctx.JSON(200, gin.H{
				"success": "project created succesfully",
				"Project": getProject,
			})
		} else {
			log.Println(err, getProject)
			ctx.JSON(400, gin.H{
				"error": "error in getting project",
			})
		}
	}
}

func DeleteProject(srvMananger *ServiceManager, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		projectId := ctx.Param("projectId")
		id, err := strconv.Atoi(projectId)
		if err != nil {
			// ... handle error
			panic(err)
		}

		projectService := &ProjectServiceManager{
			ServiceManager: srvMananger,
		}

		err, getProject := projectService.DeleteProject(id)
		if err == nil {
			log.Println("project deleted successfully", getProject)
			ctx.JSON(200, gin.H{
				"success": "project deleted succesfully",
				"Project": getProject,
			})
		} else {
			log.Println(err, getProject)
			ctx.JSON(400, gin.H{
				"error": "error in deleting project",
			})
		}
	}
}
