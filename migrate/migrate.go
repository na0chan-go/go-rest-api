package main

import (
	"log"

	"github.com/na0chan-go/go-rest-api/db"
	"github.com/na0chan-go/go-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer log.Print("Successfully migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
