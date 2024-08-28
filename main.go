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
	tr := repository.NewTaskRepository(db)
	uu := usecase.NewUserUsecase(ur)
	tu := usecase.NewTaskUsecase(tr)
	uc := controller.NewUserController(uu)
	tc := controller.NewTaskController(tu)
	e := router.NewRouter(uc, tc)
	e.Logger.Fatal(e.Start(":8080"))
}
