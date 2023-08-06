package routes

import (
	. "clokify/controllers"

	. "clokify/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(incommingRoutes *gin.Engine, db *gorm.DB) {
	srvMananger := &ServiceManager{
		Db: db,
	}

	// GET
	incommingRoutes.GET("user/:userId", GetUser(srvMananger, db))
	incommingRoutes.GET("task/:taskId", GetTask(srvMananger, db))
	incommingRoutes.GET("project/:projectId", GetProject(srvMananger, db))

	// POST
	incommingRoutes.POST("project/create", CreateProject(srvMananger, db))
	incommingRoutes.POST("task/create", CreateTask(srvMananger, db))

	// DELETE
	incommingRoutes.DELETE("task/delete", DeleteTask(srvMananger, db))
	incommingRoutes.DELETE("project/delete", DeleteProject(srvMananger, db))
}
