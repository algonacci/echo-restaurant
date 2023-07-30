package database

import (
	"github.com/algonacci/echo-restaurant/internal/model"
	"github.com/algonacci/echo-restaurant/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {

	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{}, &model.User{})

	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     40000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     50000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Nasi Goreng",
			OrderCode: "nasi_goreng",
			Price:     30000,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:      "Es Teh Manis",
			OrderCode: "es_teh_manis",
			Price:     10000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "es_jeruk",
			Price:     12000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Thai Tea",
			OrderCode: "thai_tea",
			Price:     12000,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

}
