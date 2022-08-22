package usecase

import (
	"errors"
	"log"
	"middleman-capstone/domain"
	_data "middleman-capstone/feature/orders/data"
	_helper "middleman-capstone/helper"
	"strconv"

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

func (oc *orderUseCase) GetAllAdmin(limit, offset int, role string) (data []domain.Order, err error) {

	if role != "admin" {
		log.Println("you dont have access")
		return []domain.Order{}, nil
	}
	data, err = oc.orderData.SelectDataAdminAll(limit, offset)
	return data, err
}

func (oc *orderUseCase) GetAllUser(limit, offset, idUser int) (data []domain.Order, err error) {

	if idUser == 0 {
		log.Println("order data from user not found")
		return []domain.Order{}, nil
	}
	data, err = oc.orderData.SelectDataUserAll(limit, offset, idUser)

	if err != nil {
		log.Println("failed to get data")
		return []domain.Order{}, nil
	}

	return data, err
}
func (oc *orderUseCase) GetDetail(idUser int, orderName string) (grandTotal, idOrder int, err error) {
	grandTotal, idOrder, err = oc.orderData.GetDetailData(idUser, orderName)
	if grandTotal == 0 {
		log.Println("error get data")
		return -1, 0, nil
	}
	if err != nil {
		log.Println("failed to get data")
		return 400, 0, nil
	}
	return grandTotal, idOrder, nil
}
func (oc *orderUseCase) GetItems(idOrder int) (data []domain.Items, err error) {
	data, err = oc.orderData.GetDetailItems(idOrder)

	if err != nil {
		log.Println("failed to get data")
		return []domain.Items{}, nil
	}
	return data, nil
}

func (oc *orderUseCase) GetIncoming(limit, offset int, role string) (data []domain.Order, err error) {

	if role != "admin" {
		log.Println("you dont have access")
		return []domain.Order{}, errors.New("you dont have access")
	}
	data, err = oc.orderData.GetIncomingData(limit, offset)

	if err != nil {
		log.Println("failed to get data")
		return []domain.Order{}, nil
	}
	return data, nil
}

func (oc *orderUseCase) CreateOrder(dataOrder domain.Order, idUser int) int {

	validError := oc.validate.Var(dataOrder, "gt=0")
	if validError != nil {
		log.Println("Validation error : ", validError)
		return 400
	}

	idOrder, err := oc.orderData.Insert(dataOrder)
	if err != nil {
		log.Println("error adding data")
		return 400
	}

	create := oc.orderData.InsertData(dataOrder.Items, idOrder)
	if len(create) == 0 {
		log.Println("error after creating data")
		return 500
	}

	delete := oc.orderData.DeleteCart(idUser)
	if !delete {
		log.Println("failed delete data")
		return 500
	}

	return 201
}

func (oc *orderUseCase) Payment(grandTotal, idUser int) (orderName, url, token string, dataUser domain.User) {
	user, _ := oc.orderData.GetUser(idUser)
	if user.ID == 0 {
		log.Println("failed to get user data")
		return "", "", "", domain.User{}
	}
	data := _data.OrderPayment{
		Email:      user.Email,
		Name:       user.Name,
		GrandTotal: grandTotal,
	}
	data.Phone, _ = strconv.Atoi(user.Phone)

	orderIDGen, trans := _helper.Payment(data)
	if orderIDGen == "" {
		log.Println("failed to make midtrans invoice")
	}

	return orderIDGen, trans.RedirectURL, trans.Token, user
}

func (oc *orderUseCase) AcceptPayment(data domain.PaymentWeb) (row int, err error) {

	if data.TransactionStatus == "settlement" {
		_, err := oc.orderData.AcceptPaymentData(data)
		if err != nil {
			log.Println("failed to update status")
			return 500, err
		}
	} else if data.TransactionStatus == "expire" {
		_, err := oc.orderData.CancelPaymentData(data)
		if err != nil {
			log.Println("failed to update status")
			return 500, err
		}
	} else if data.TransactionStatus == "cancel" {
		_, err := oc.orderData.CancelPaymentData(data)
		if err != nil {
			log.Println("failed to update status")
			return 500, err
		}
	}

	if data.TransactionStatus == "" {
		log.Println("error payment")
		return -1, err
	}
	return row, err
}

func (oc *orderUseCase) ConfirmOrder(orderName, role string) (domain.Order, int) {
	if role != "admin" {
		log.Println("you dont have access")
		return domain.Order{}, 401
	}

	order := oc.orderData.ConfirmOrderData(orderName)
	if order.ID == 0 {
		log.Println("data not found")
		return domain.Order{}, 404
	}

	user, _ := oc.orderData.GetUser(order.UserID)
	if user.ID == 0 {
		log.Println("failed to get user data")
		return domain.Order{}, 500
	}

	totalPayment := strconv.Itoa(order.GrandTotal)
	data := _helper.Recipient{
		OrderID:      orderName,
		Name:         user.Name,
		Email:        user.Email,
		Handphone:    user.Phone,
		TotalPayment: totalPayment,
	}

	if data.OrderID == "" {
		log.Println("Empty Data")
		return domain.Order{}, 404
	} else {
		_helper.SendEmail(data)
	}

	return order, 200
}

func (oc *orderUseCase) DoneOrder(ordername string) (domain.Order, int) {
	order := oc.orderData.DoneOrderData(ordername)
	if order.ID == 0 {
		log.Println("Empty Data")
		return domain.Order{}, 404
	}

	updateadminstok := oc.orderData.UpdateStokAdmin(ordername)
	if !updateadminstok {
		log.Println("failed update data")
		return domain.Order{}, 500
	}

	cekown, id := oc.orderData.CekUser(ordername)
	if len(cekown) < 1 {
		log.Println("failed retrieve data")
		return domain.Order{}, 500
	}

	for _, val := range cekown {
		owned := oc.orderData.CekOwned(val, id)
		if !owned {
			product := oc.orderData.CreateNewProduct(val, id)
			if !product {
				log.Println("failed create data")
				return domain.Order{}, 500
			}
		} else {
			product := oc.orderData.UpdateNewProduct(val, id)
			if !product {
				log.Println("failed update data")
				return domain.Order{}, 500
			}
		}
	}

	return order, 200
}
