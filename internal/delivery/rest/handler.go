package rest

import "github.com/algonacci/echo-restaurant/internal/usecase/resto"

type handler struct {
	restoUsecase resto.Usecase
}

func NewHandler(restoUsecase resto.Usecase) *handler {
	return &handler{
		restoUsecase: restoUsecase,
	}
}
