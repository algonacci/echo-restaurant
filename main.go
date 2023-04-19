package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=admin dbname=go_resto_app sslmode=disable"
)

type MenuType string

const (
	MenuTypeFood  = "food"
	MenuTypeDrink = "drink"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      string
}

func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     40000,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     50000,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Nasi Goreng",
			OrderCode: "nasi_goreng",
			Price:     30000,
			Type:      MenuTypeFood,
		},
	}

	drinkMenu := []MenuItem{
		{
			Name:      "Es Teh Manis",
			OrderCode: "es_teh_manis",
			Price:     10000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "es_jeruk",
			Price:     12000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Thai Tea",
			OrderCode: "thai_tea",
			Price:     12000,
			Type:      MenuTypeDrink,
		},
	}

	fmt.Println(foodMenu, drinkMenu)

	// db, err := gorm.Open(postgres.Open(dbAddress))
	// if err != nil {
	// 	panic(err)
	// }
	// db.AutoMigrate(&MenuItem{})

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}
	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

}

func getFoodMenu(c echo.Context) error {
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}
	var menuData []MenuItem
	db.Where(MenuItem{Type: MenuTypeFood}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func getDrinkMenu(c echo.Context) error {
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}
	var menuData []MenuItem
	db.Where(MenuItem{Type: MenuTypeDrink}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}
	var menuData []MenuItem
	db.Where(MenuItem{Type: menuType}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func main() {
	seedDB()
	e := echo.New()
	e.GET("/menu", getMenu)
	e.Logger.Fatal(e.Start(":8080"))
}
