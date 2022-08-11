package usecase

import (
	"errors"
	"middleman-capstone/domain"
)

type cartUseCase struct {
	cartData domain.ChartData
}

func New(pd domain.ChartData) domain.CartUseCase {
	return &cartUseCase{
		cartData: pd,
	}
}
func (uc *cartUseCase) GetAllData(limit, offset, idFromToken int) (data []domain.Cart, err error) {
	data, err = uc.cartData.SelectData(limit, offset, idFromToken)
	for k, v := range data {
		data[k].TotalPrice = v.Qty * v.Product.Price
	}
	return data, err
}

func (uc *cartUseCase) CreateData(data domain.Cart) (row int, err error) {
	if data.Qty == 0 || data.Product.ID == 0 {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	isExist, idCart, Qty, _ := uc.cartData.CheckCart(data.Product.ID, data.UserID)
	if isExist {
		row, err = uc.cartData.UpdateDataDB(Qty+1, idCart, data.UserID)
	} else {
		row, err = uc.cartData.InsertData(data)
	}

	return row, err
}

func (uc *cartUseCase) UpdateData(qty, idCart, idFromToken int) (row int, err error) {
	row, err = uc.cartData.UpdateDataDB(qty, idCart, idFromToken)
	return row, err
}
