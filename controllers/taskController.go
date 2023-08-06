package controllers

import (
	. "clokify/services"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"

	. "clokify/models"
	. "clokify/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTask(srvMananger *ServiceManager, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		taskId := ctx.Param("taskId")
		id, err := strconv.Atoi(taskId)
		if err != nil {
			// ... handle error
			panic(err)
		}

		taskService := &TaskServiceManager{
			ServiceManager: srvMananger,
		}

		err, getTask := taskService.GetTask(id)
		if err == nil {
			log.Println("task fetched successfully", getTask)
			ctx.JSON(200, gin.H{
				"success": "task fetched succesfully",
				"Task":    getTask,
			})
			return
		}

		log.Println(err, getTask)
		ctx.JSON(400, gin.H{
			"error": "error in getting task",
		})
	}
}

func DeleteTask(srvMananger *ServiceManager, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		taskId := ctx.Param("taskId")
		id, err := strconv.Atoi(taskId)
		if err != nil {
			// ... handle error
			panic(err)
		}

		taskService := &TaskServiceManager{
			ServiceManager: srvMananger,
		}

		err, getTask := taskService.DeleteTask(id)
		if err == nil {
			log.Println("task deleted successfully", getTask)
			ctx.JSON(200, gin.H{
				"success": "task deleted succesfully",
				"Task":    getTask,
			})
			return
		}

		log.Println(err, getTask)
		ctx.JSON(400, gin.H{
			"error": "error in deleting task",
		})
	}
}

func CreateTask(srvMananger *ServiceManager, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var task TaskCreateType
		taskService := &TaskServiceManager{
			ServiceManager: srvMananger,
		}

		bodyData, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Failed to read request body",
			})
			return
		}

		if err := json.Unmarshal(bodyData, &task); err != nil {
			fmt.Print(task)
			ctx.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		err, projectRes := taskService.CreateTask(&task, srvMananger)
		if err == nil {
			log.Println("task created succesfully", projectRes)
			ctx.JSON(200, gin.H{"success": "task created succesfully"})
			return
		}

		ctx.JSON(400, gin.H{
			"error":  "task didnt created succesfully",
			"reason": err.Error(),
		})
	}
}
