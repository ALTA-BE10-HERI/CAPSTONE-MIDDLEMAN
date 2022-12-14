// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "middleman-capstone/domain"

	mock "github.com/stretchr/testify/mock"
)

// InventoryUseCase is an autogenerated mock type for the InventoryUseCase type
type InventoryUseCase struct {
	mock.Mock
}

// CreateAdminInventory provides a mock function with given fields: newRecap, id, role
func (_m *InventoryUseCase) CreateAdminInventory(newRecap domain.Inventory, id int, role string) (domain.Inventory, int) {
	ret := _m.Called(newRecap, id, role)

	var r0 domain.Inventory
	if rf, ok := ret.Get(0).(func(domain.Inventory, int, string) domain.Inventory); ok {
		r0 = rf(newRecap, id, role)
	} else {
		r0 = ret.Get(0).(domain.Inventory)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(domain.Inventory, int, string) int); ok {
		r1 = rf(newRecap, id, role)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// CreateUserInventory provides a mock function with given fields: newRecap, idUser
func (_m *InventoryUseCase) CreateUserInventory(newRecap domain.Inventory, idUser int) (domain.Inventory, int) {
	ret := _m.Called(newRecap, idUser)

	var r0 domain.Inventory
	if rf, ok := ret.Get(0).(func(domain.Inventory, int) domain.Inventory); ok {
		r0 = rf(newRecap, idUser)
	} else {
		r0 = ret.Get(0).(domain.Inventory)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(domain.Inventory, int) int); ok {
		r1 = rf(newRecap, idUser)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// ReadAdminOutBoundDetail provides a mock function with given fields: outboundIDGenerate
func (_m *InventoryUseCase) ReadAdminOutBoundDetail(outboundIDGenerate string) ([]domain.InventoryProduct, int, string) {
	ret := _m.Called(outboundIDGenerate)

	var r0 []domain.InventoryProduct
	if rf, ok := ret.Get(0).(func(string) []domain.InventoryProduct); ok {
		r0 = rf(outboundIDGenerate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.InventoryProduct)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(outboundIDGenerate)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(string) string); ok {
		r2 = rf(outboundIDGenerate)
	} else {
		r2 = ret.Get(2).(string)
	}

	return r0, r1, r2
}

// ReadAdminOutBoundHistory provides a mock function with given fields:
func (_m *InventoryUseCase) ReadAdminOutBoundHistory() ([]domain.Inventory, int) {
	ret := _m.Called()

	var r0 []domain.Inventory
	if rf, ok := ret.Get(0).(func() []domain.Inventory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Inventory)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// ReadUserOutBoundDetail provides a mock function with given fields: id, outboundIDGenerate
func (_m *InventoryUseCase) ReadUserOutBoundDetail(id int, outboundIDGenerate string) ([]domain.InventoryProduct, int, string) {
	ret := _m.Called(id, outboundIDGenerate)

	var r0 []domain.InventoryProduct
	if rf, ok := ret.Get(0).(func(int, string) []domain.InventoryProduct); ok {
		r0 = rf(id, outboundIDGenerate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.InventoryProduct)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, string) int); ok {
		r1 = rf(id, outboundIDGenerate)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(int, string) string); ok {
		r2 = rf(id, outboundIDGenerate)
	} else {
		r2 = ret.Get(2).(string)
	}

	return r0, r1, r2
}

// ReadUserOutBoundHistory provides a mock function with given fields: id
func (_m *InventoryUseCase) ReadUserOutBoundHistory(id int) ([]domain.Inventory, int) {
	ret := _m.Called(id)

	var r0 []domain.Inventory
	if rf, ok := ret.Get(0).(func(int) []domain.Inventory); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Inventory)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int) int); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

type mockConstructorTestingTNewInventoryUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewInventoryUseCase creates a new instance of InventoryUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInventoryUseCase(t mockConstructorTestingTNewInventoryUseCase) *InventoryUseCase {
	mock := &InventoryUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
