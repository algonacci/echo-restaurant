package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     40000,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     50000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

func getDrinkMenu(c echo.Context) error {
	drinkMenu := []MenuItem{
		{
			Name:      "Es Teh Manis",
			OrderCode: "es_teh_manis",
			Price:     10000,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "es_jeruk",
			Price:     12000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinkMenu,
	})
}

func main() {
	e := echo.New()
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":8080"))
}
