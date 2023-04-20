package resto

import "github.com/algonacci/echo-restaurant/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
