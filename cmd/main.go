package main

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/algonacci/echo-restaurant/internal/database"
	"github.com/algonacci/echo-restaurant/internal/delivery/rest"
	mRepo "github.com/algonacci/echo-restaurant/internal/repository/menu"
	oRepo "github.com/algonacci/echo-restaurant/internal/repository/order"
	uRepo "github.com/algonacci/echo-restaurant/internal/repository/user"
	rUsecase "github.com/algonacci/echo-restaurant/internal/usecase/resto"
	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=admin dbname=go_resto_app sslmode=disable"
)

func main() {
	e := echo.New()

	db := database.GetDB(dbAddress)
	secret := "AES256Key-32Characters1234567890"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64&1024, 4, 32, signKey, 60*time.Second)
	if err != nil {
		panic(err)
	}

	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo, userRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":8080"))
}
