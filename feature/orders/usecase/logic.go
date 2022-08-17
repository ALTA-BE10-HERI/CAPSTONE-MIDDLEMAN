package usecase

import (
	"log"
	"middleman-capstone/domain"

	"github.com/go-playground/validator"
)

type orderUseCase struct {
	orderData domain.OrderData
	validate  *validator.Validate
}

func New(od domain.OrderData, v *validator.Validate) domain.OrderUseCase {
	return &orderUseCase{
		orderData: od,
		validate:  v,
	}
}

// func (oc *orderUseCase) GetAllAdmin(limit, offset int) (data []domain.Order, err error) {
// 	data, err = oc.orderData.SelectDataAdminAll(limit, offset)
// 	return data, err
// }

func (oc *orderUseCase) CreateOrder(dataOrder domain.Order, id int) int {

	validError := oc.validate.Var(dataOrder, "gt=0")
	if validError != nil {
		log.Println("Validation error : ", validError)
		return 400
	}

	create := oc.orderData.InsertData(dataOrder.Items, id)
	if len(create) == 0 {
		log.Println("error after creating data")
		return 500
	}
	// fmt.Println("dataOrder :", dataOrder)
	// fmt.Println("dataOrder :", id)
	// fmt.Println("error :", err)
	// fmt.Println("error2 :", err3)
	return 201
}

// func (oc *orderUseCase) CreateItems(data []domain.Items) (row int, err error) {
// 	res, err := oc.orderData.CreateItems(data, 1)
// 	return res, err
// }
