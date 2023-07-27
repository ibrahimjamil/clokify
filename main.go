package main

import (
	. "clokify/config"
	. "clokify/db"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed to load env variables")
	}

	db, err := DbConnection(EnvConfig())
	if err != nil {
		log.Fatal("Could not load the database")
	}

	fmt.Printf(db.Name())
}
