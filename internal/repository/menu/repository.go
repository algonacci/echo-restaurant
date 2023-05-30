package menu

import (
	"github.com/algonacci/echo-restaurant/internal/model"
)

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) (model.MenuItem, error)
}
