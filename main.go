package main

import (
	. "clokify/db"
	routes "clokify/routes"
	"fmt"
	"log"

	. "clokify/config"

	. "clokify/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// loading env
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Failed to load env variables")
	}

	// db connection
	db, err := DbConnection()
	if err != nil {
		log.Panic("Could not initialize the database")
	}

	// assign port number
	fmt.Print(EnvConfig().AppPort, EnvConfig().DBName, EnvConfig().Host)
	AppPort := EnvConfig().AppPort
	if AppPort == "" {
		AppPort = "8000"
	}

	// making router with gin package
	router := gin.New()
	router.Use(gin.Logger())

	// un auth and auth routes
	routes.UnAuthRoutes(router, db)
	router.Use(AuthMiddleware())
	routes.AuthRoutes(router, db)

	// app running on port
	router.Run(":" + AppPort)
}
