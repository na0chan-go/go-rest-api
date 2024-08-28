package main

import (
	"log"

	"github.com/na0chan-go/go-rest-api/controller"
	"github.com/na0chan-go/go-rest-api/db"
	"github.com/na0chan-go/go-rest-api/repository"
	"github.com/na0chan-go/go-rest-api/router"
	"github.com/na0chan-go/go-rest-api/usecase"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatalf("Error creating db: %v", err)
	}
	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)
	e := router.NewRouter(uc)
	e.Logger.Fatal(e.Start(":8080"))
}
