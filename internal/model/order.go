package model

import "github.com/algonacci/echo-restaurant/internal/model/constant"

type Order struct {
	ID            string               `gorm:"primaryKey" json:"id"`
	Status        constant.OrderStatus `json:"status"`
	ProductOrders []ProductOrder       `json:"product_orders"`
	ReferenceID   string               `gorm:"unique" json:"reference_id"`
}

type ProductOrder struct {
	ID         string
	OrderID    string
	OrderCode  string
	Quantity   int
	TotalPrice int64
	Status     constant.ProductOrderStatus
}

type OrderMenuProductRequest struct {
	OrderCode string `json:"order_code"`
	Quantity  int    `json:"quantity"`
}

type OrderMenuRequest struct {
	OrderProducts []OrderMenuProductRequest `json:"order_products"`
	ReferenceID   string                    `json:"reference_id"`
}

type GetOrderInfoRequest struct {
	OrderID string
}
