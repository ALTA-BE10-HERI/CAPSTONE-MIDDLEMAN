package usecase

import (
	"fmt"
	"middleman-capstone/domain"
)

type orderUseCase struct {
	orderData domain.OrderData
}

func New(od domain.OrderData) domain.OrderUseCase {
	return &orderUseCase{
		orderData: od,
	}
}

func (oc *orderUseCase) GetAllAdmin(limit, offset int) (data []domain.Order, err error) {
	data, err = oc.orderData.SelectDataAdminAll(limit, offset)
	return data, err
}

func (oc *orderUseCase) CreateOrder(dataOrder domain.Order) (row int, err error) {
	id, err3 := oc.orderData.InsertData(dataOrder)
	fmt.Println("dataOrder :", dataOrder)
	fmt.Println("dataOrder :", id)
	fmt.Println("error :", err)
	fmt.Println("error2 :", err3)
	return id, err3
}

func (oc *orderUseCase) CreateItems(data []domain.Items) (row int, err error) {
	res, err := oc.orderData.CreateItems(data, 1)
	return res, err
}
