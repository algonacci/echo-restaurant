package resto

import (
	"errors"

	"github.com/algonacci/echo-restaurant/internal/model"
	"github.com/algonacci/echo-restaurant/internal/model/constant"
	"github.com/algonacci/echo-restaurant/internal/repository/menu"
	"github.com/algonacci/echo-restaurant/internal/repository/order"
	"github.com/algonacci/echo-restaurant/internal/repository/user"
	"github.com/google/uuid"
)

type restoUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
	userRepo  user.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository, userRepo user.Repository) Usecase {
	return &restoUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
		userRepo:  userRepo,
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

func (r *restoUsecase) GetOrderInfo(request model.GetOrderInfoRequest) (model.Order, error) {
	orderData, err := r.orderRepo.GetOrderInfo(request.OrderID)
	if err != nil {
		return orderData, err
	}

	return orderData, nil
}

func (r *restoUsecase) RegisterUser(request model.RegisterRequest) (model.User, error) {
	userRegistered, err := r.userRepo.CheckRegistered(request.Username)
	if err != nil {
		return model.User{}, err
	}
	if userRegistered {
		return model.User{}, errors.New("user already registered")
	}

	userHash, err := r.userRepo.GenerateUserHash(request.Password)
	if err != nil {
		return model.User{}, err
	}

	userData, err := r.userRepo.RegisterUser(model.User{
		ID:       uuid.New().String(),
		Username: request.Username,
		Hash:     userHash,
	})

	if err != nil {
		return model.User{}, err
	}

	return userData, nil
}
