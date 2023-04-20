package main

import (
	"github.com/algonacci/echo-restaurant/internal/database"
	"github.com/algonacci/echo-restaurant/internal/delivery/rest"
	mRepo "github.com/algonacci/echo-restaurant/internal/repository/menu"
	rUseCase "github.com/algonacci/echo-restaurant/internal/usecase/resto"
	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=admin dbname=go_resto_app sslmode=disable"
)

func main() {
	e := echo.New()
	db := database.GetDB(dbAddress)
	menuRepo := mRepo.GetRepository(db)
	restoUsecase := rUseCase.GetUseCase(menuRepo)
	h := rest.NewHandler(restoUsecase)
	rest.LoadRoutes(e, h)
	e.Logger.Fatal(e.Start(":8080"))
}
