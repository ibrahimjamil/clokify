package main

import (
	. "clokify/db"
	. "clokify/types"
	. "clokify/utils/cruds"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Failed to load env variables")
	}

	db, err := DbConnection()
	if err != nil {
		log.Panic("Could not initialize the database")
	}

	// init global service manager each service will extend that
	srvMananger := &ServiceManager{
		Db: db,
	}

	// can uncomment delete user to check project creation scenario
	UserCrud(db, srvMananger)
	ProjectCrud(db, srvMananger)
	TaskCrud(db, srvMananger)
}
