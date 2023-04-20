package menu

import (
	"github.com/algonacci/echo-restaurant/internal/model"
)

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
