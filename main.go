package main

import (
	"log"

	"github.com/na0chan-go/go-rest-api/controller"
	"github.com/na0chan-go/go-rest-api/db"
	"github.com/na0chan-go/go-rest-api/repository"
	"github.com/na0chan-go/go-rest-api/router"
	"github.com/na0chan-go/go-rest-api/usecase"
	"github.com/na0chan-go/go-rest-api/validator"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatalf("Error creating db: %v", err)
	}
	uv := validator.NewUserValidator()
	tv := validator.NewTaskValidator()
	ur := repository.NewUserRepository(db)
	tr := repository.NewTaskRepository(db)
	uu := usecase.NewUserUsecase(ur, uv)
	tu := usecase.NewTaskUsecase(tr, tv)
	uc := controller.NewUserController(uu)
	tc := controller.NewTaskController(tu)
	e := router.NewRouter(uc, tc)
	e.Logger.Fatal(e.Start(":8080"))
}
