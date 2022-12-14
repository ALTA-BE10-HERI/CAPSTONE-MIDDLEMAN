// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "middleman-capstone/domain"

	mock "github.com/stretchr/testify/mock"
)

// OrderData is an autogenerated mock type for the OrderData type
type OrderData struct {
	mock.Mock
}

// AcceptPaymentData provides a mock function with given fields: data
func (_m *OrderData) AcceptPaymentData(data domain.PaymentWeb) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.PaymentWeb) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.PaymentWeb) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelPaymentData provides a mock function with given fields: data
func (_m *OrderData) CancelPaymentData(data domain.PaymentWeb) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.PaymentWeb) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.PaymentWeb) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CekOwned provides a mock function with given fields: product, userid
func (_m *OrderData) CekOwned(product domain.Items, userid int) bool {
	ret := _m.Called(product, userid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.Items, int) bool); ok {
		r0 = rf(product, userid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CekUser provides a mock function with given fields: ordername
func (_m *OrderData) CekUser(ordername string) ([]domain.Items, int) {
	ret := _m.Called(ordername)

	var r0 []domain.Items
	if rf, ok := ret.Get(0).(func(string) []domain.Items); ok {
		r0 = rf(ordername)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Items)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(ordername)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// ConfirmOrderData provides a mock function with given fields: orderid
func (_m *OrderData) ConfirmOrderData(orderid string) domain.Order {
	ret := _m.Called(orderid)

	var r0 domain.Order
	if rf, ok := ret.Get(0).(func(string) domain.Order); ok {
		r0 = rf(orderid)
	} else {
		r0 = ret.Get(0).(domain.Order)
	}

	return r0
}

// CreateNewProduct provides a mock function with given fields: product, userid
func (_m *OrderData) CreateNewProduct(product domain.Items, userid int) bool {
	ret := _m.Called(product, userid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.Items, int) bool); ok {
		r0 = rf(product, userid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DeleteCart provides a mock function with given fields: userid
func (_m *OrderData) DeleteCart(userid int) bool {
	ret := _m.Called(userid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(userid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DoneOrderData provides a mock function with given fields: orderid
func (_m *OrderData) DoneOrderData(orderid string) domain.Order {
	ret := _m.Called(orderid)

	var r0 domain.Order
	if rf, ok := ret.Get(0).(func(string) domain.Order); ok {
		r0 = rf(orderid)
	} else {
		r0 = ret.Get(0).(domain.Order)
	}

	return r0
}

// GetDetailData provides a mock function with given fields: orderName
func (_m *OrderData) GetDetailData(orderName string) (int, int, string, error) {
	ret := _m.Called(orderName)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(orderName)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(orderName)
	} else {
		r1 = ret.Get(1).(int)
	}
	
	var r2 string
	if rf, ok := ret.Get(2).(func(string) string); ok {
		r2 = rf(orderName)
	} else {
		r2 = ret.Get(2).(string)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(string) error); ok {
		r3 = rf(orderName)
	} else {
		r3 = ret.Error(3)
	}


	return r0, r1, r2, r3
}

// GetDetailItems provides a mock function with given fields: idOrder
func (_m *OrderData) GetDetailItems(idOrder int) ([]domain.Items, error) {
	ret := _m.Called(idOrder)

	var r0 []domain.Items
	if rf, ok := ret.Get(0).(func(int) []domain.Items); ok {
		r0 = rf(idOrder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Items)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idOrder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIncomingData provides a mock function with given fields: limit, offset
func (_m *OrderData) GetIncomingData(limit int, offset int) ([]domain.Order, error) {
	ret := _m.Called(limit, offset)

	var r0 []domain.Order
	if rf, ok := ret.Get(0).(func(int, int) []domain.Order); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: idUser
func (_m *OrderData) GetUser(idUser int) (domain.User, error) {
	ret := _m.Called(idUser)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(int) domain.User); ok {
		r0 = rf(idUser)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByOrderName provides a mock function with given fields: orderName
func (_m *OrderData) GetUserByOrderName(orderName string) (domain.Order, error) {
	ret := _m.Called(orderName)

	var r0 domain.Order
	if rf, ok := ret.Get(0).(func(string) domain.Order); ok {
		r0 = rf(orderName)
	} else {
		r0 = ret.Get(0).(domain.Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orderName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: data
func (_m *OrderData) Insert(data domain.Order) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.Order) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Order) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertData provides a mock function with given fields: data, id
func (_m *OrderData) InsertData(data []domain.Items, id int) []domain.Items {
	ret := _m.Called(data, id)

	var r0 []domain.Items
	if rf, ok := ret.Get(0).(func([]domain.Items, int) []domain.Items); ok {
		r0 = rf(data, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Items)
		}
	}

	return r0
}

// SelectDataAdminAll provides a mock function with given fields: limit, offset
func (_m *OrderData) SelectDataAdminAll(limit int, offset int) ([]domain.Order, error) {
	ret := _m.Called(limit, offset)

	var r0 []domain.Order
	if rf, ok := ret.Get(0).(func(int, int) []domain.Order); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectDataUserAll provides a mock function with given fields: limit, offset, idUser
func (_m *OrderData) SelectDataUserAll(limit int, offset int, idUser int) ([]domain.Order, error) {
	ret := _m.Called(limit, offset, idUser)

	var r0 []domain.Order
	if rf, ok := ret.Get(0).(func(int, int, int) []domain.Order); ok {
		r0 = rf(limit, offset, idUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, int) error); ok {
		r1 = rf(limit, offset, idUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNewProduct provides a mock function with given fields: product, userid
func (_m *OrderData) UpdateNewProduct(product domain.Items, userid int) bool {
	ret := _m.Called(product, userid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.Items, int) bool); ok {
		r0 = rf(product, userid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UpdateStokAdmin provides a mock function with given fields: ordername
func (_m *OrderData) UpdateStokAdmin(ordername string) bool {
	ret := _m.Called(ordername)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(ordername)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewOrderData interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderData creates a new instance of OrderData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderData(t mockConstructorTestingTNewOrderData) *OrderData {
	mock := &OrderData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
