package routes

import (
	. "clokify/controllers"

	. "clokify/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UnAuthRoutes(incommingRoutes *gin.Engine, db *gorm.DB) {
	srvMananger := &ServiceManager{
		Db: db,
	}

	// POST
	incommingRoutes.POST("user/register", Register(srvMananger, db))
	incommingRoutes.POST("user/login", Login(srvMananger, db))

	// GET
	incommingRoutes.GET("check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": "OKk",
		})
	})
}
