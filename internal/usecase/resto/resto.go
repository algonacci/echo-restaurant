package resto

import (
	"github.com/algonacci/echo-restaurant/internal/model"
	"github.com/algonacci/echo-restaurant/internal/model/constant"
	"github.com/algonacci/echo-restaurant/internal/repository/menu"
	"github.com/algonacci/echo-restaurant/internal/repository/order"
	"github.com/google/uuid"
)

type restoUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository) Usecase {
	return &restoUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
	}
}

func (m *restoUsecase) GetMenuList(menuType string) ([]model.MenuItem, error) {
	return m.menuRepo.GetMenuList(menuType)
}

func (m *restoUsecase) Order(request model.OrderMenuRequest) (model.Order, error) {
	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menuData, err := m.menuRepo.GetMenu(orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}

		productOrderData[i] = model.ProductOrder{
			ID:         uuid.New().String(),
			OrderCode:  orderProduct.OrderCode,
			Quantity:   orderProduct.Quantity,
			TotalPrice: int64(menuData.Price) * int64(orderProduct.Quantity),
			Status:     constant.ProductOrderStatusPreparing,
		}
	}

	orderData := model.Order{
		ID:            uuid.New().String(),
		Status:        constant.OrderStatusProcessed,
		ProductOrders: productOrderData,
		ReferenceID:   request.ReferenceID,
	}

	createdOrderData, err := m.orderRepo.CreateOrder(orderData)
	if err != nil {
		return model.Order{}, err
	}

	return createdOrderData, nil
}

func (m *restoUsecase) GetOrderInfo(request model.GetOrderInfoRequest) (model.Order, error) {
	orderData, err := m.orderRepo.GetOrderInfo(request.OrderID)
	if err != nil {
		return orderData, err
	}

	return orderData, nil
}
