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

func Register(srvMananger *ServiceManager, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user User
		userService := &UserServiceManager{
			ServiceManager: srvMananger,
		}

		bodyData, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Failed to read request body",
			})
			return
		}

		if err := json.Unmarshal(bodyData, &user); err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		err, userRes := userService.CreateUser(&user)
		if err == nil {
			log.Println("user created successfully", userRes)
			ctx.JSON(200, gin.H{"success": "user registered successfully"})
			return
		}

		log.Println(err)
		ctx.JSON(400, gin.H{"success": "user didnt registered successfully some issue in creating user"})
	}
}

func Login(srvMananger *ServiceManager, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginType LoginType
		userService := &UserServiceManager{
			ServiceManager: srvMananger,
		}

		bodyData, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "Failed to read request body",
			})
			return
		}

		if err := json.Unmarshal(bodyData, &loginType); err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		err, userRes, bool := userService.LoginUser(&loginType)
		if err == nil {
			log.Println("user login successfully", userRes, bool)
			ctx.JSON(200, gin.H{
				"success": "user login successfully",
				"token":   userRes,
			})
			return
		}

		log.Println(err)
		ctx.JSON(400, gin.H{"error": "user didnt login successfully some issue in login user"})
	}
}

func GetUser(srvMananger *ServiceManager, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("userId")
		id, err := strconv.Atoi(userId)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "some error in userId string to int conversion",
			})
		}

		userService := &UserServiceManager{
			ServiceManager: srvMananger,
		}

		err, getUser := userService.GetUser(id)
		if err == nil {
			log.Println("user fetched successfully", getUser)
			ctx.JSON(200, gin.H{
				"success": "user fetched succesfully",
				"Project": getUser,
			})
			return
		}

		ctx.JSON(400, gin.H{
			"error": "error in getting user",
		})
	}
}
