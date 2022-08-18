package usecase

import (
	"errors"
	"log"
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
func (uc *cartUseCase) GetAllData(limit, offset, idFromToken int) (data []domain.Cart, total int, err error) {
	data, err = uc.cartData.SelectData(limit, offset, idFromToken)
	subTotal := 0
	for _, v := range data {
		subTotal += v.Subtotal
	}
	return data, subTotal, err
}

func (uc *cartUseCase) CreateData(data domain.Cart) (row int, err error) {
	if data.Qty == 0 || data.Product.ID == 0 {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	if data.ID == 0 {
		return 404, errors.New("data product not found")
	}
	productStock, _ := uc.cartData.GetStockProduct(data.Product.ID)
	// cekStock := data.Qty >= productStock
	log.Println("qty : ", data.Qty, "stok :", productStock)
	if data.Qty > productStock {
		return 400, errors.New("qty exceeds product stock")
	}
	productPrice, _ := uc.cartData.GetPriceProduct(data.Product.ID)
	isExist, idCart, Qty, _ := uc.cartData.CheckCart(data.Product.ID, data.UserID)
	if isExist {
		cekCurrentQty, _ := uc.cartData.GetQtyProductCart(idCart)
		// log.Println("cek id :", idCart)
		if (cekCurrentQty + data.Qty) > productStock {
			return 400, errors.New("qty exceeds product stock")
		}
		data.Subtotal = productPrice * Qty
		row, err = uc.cartData.UpdateDataDB(Qty+data.Qty, idCart, data.UserID)
	} else {
		data.Subtotal = productPrice * data.Qty
		row, err = uc.cartData.InsertData(data)
	}

	return row, err
}

func (uc *cartUseCase) UpdateData(qty, idCart, idFromToken int) (row int, err error) {
	row, err = uc.cartData.UpdateDataDB(qty, idCart, idFromToken)
	return row, err
}
func (uc *cartUseCase) DeleteData(idProd, idFromToken int) (row int, err error) {
	row, err = uc.cartData.DeleteDataDB(idProd, idFromToken)
	return row, err
}
