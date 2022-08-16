package usecase

import "middleman-capstone/domain"

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
