package usecase

import (
	"middleman-capstone/domain"
	"middleman-capstone/domain/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserInventory(t *testing.T) {
	repo := new(mocks.InventoryData)
	input := []domain.InventoryProduct{{ProductID: 2, Qty: 3, Unit: "kg"}, {ProductID: 5, Qty: 7, Unit: "kg"}}
	input2 := domain.Inventory{ID: 2, OutBound: "6981234678", UserID: 5, InventoryProduct: []domain.InventoryProduct{}}
	input3 := []domain.InventoryProduct{}
	input4 := domain.Inventory{ID: 0, InventoryProduct: []domain.InventoryProduct{}}

	t.Run("insufficient amount", func(t *testing.T) {
		repo.On("CekStok", mock.Anything, mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateUserInventory(input2, 1)

		assert.Equal(t, 404, status)
		assert.Equal(t, domain.Inventory{}, bound)

	})

	t.Run("error after creating outbound data", func(t *testing.T) {
		repo.On("CekStok", mock.Anything, mock.Anything).Return(true).Once()
		repo.On("CreateUserInventoryData", mock.Anything, mock.Anything, mock.Anything).Return(input4).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateUserInventory(input2, 1)

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.Inventory{}, bound)

	})

	t.Run("error after creating data", func(t *testing.T) {
		repo.On("CekStok", mock.Anything, mock.Anything).Return(true).Once()
		repo.On("CreateUserInventoryData", mock.Anything, mock.Anything, mock.Anything).Return(input2).Once()
		repo.On("CreateUserDetailInventoryData", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(input3).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateUserInventory(input2, 1)

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.Inventory{}, bound)

	})

	t.Run("error after update data", func(t *testing.T) {
		repo.On("CekStok", mock.Anything, mock.Anything).Return(true).Once()
		repo.On("CreateUserInventoryData", mock.Anything, mock.Anything, mock.Anything).Return(input2).Once()
		repo.On("CreateUserDetailInventoryData", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(input).Once()
		repo.On("RekapStock", mock.Anything, mock.Anything, mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateUserInventory(input2, 1)

		assert.Equal(t, 404, status)
		assert.Equal(t, domain.Inventory{}, bound)

	})

	t.Run("error after update data", func(t *testing.T) {
		repo.On("CekStok", mock.Anything, mock.Anything).Return(true).Once()
		repo.On("CreateUserInventoryData", mock.Anything, mock.Anything, mock.Anything).Return(input2).Once()
		repo.On("CreateUserDetailInventoryData", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(input).Once()
		repo.On("RekapStock", mock.Anything, mock.Anything, mock.Anything).Return(true).Once()
		repo.On("DeleteInOutBound", mock.Anything).Return("cannot delete data").Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateUserInventory(input2, 1)

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.Inventory{}, bound)

	})
	t.Run("error after update data", func(t *testing.T) {
		repo.On("CekStok", mock.Anything, mock.Anything).Return(true).Once()
		repo.On("CreateUserInventoryData", mock.Anything, mock.Anything, mock.Anything).Return(input2).Once()
		repo.On("CreateUserDetailInventoryData", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(input).Once()
		repo.On("RekapStock", mock.Anything, mock.Anything, mock.Anything).Return(true).Once()
		repo.On("DeleteInOutBound", mock.Anything).Return("no data deleted").Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateUserInventory(input2, 1)

		assert.Equal(t, 404, status)
		assert.Equal(t, domain.Inventory{}, bound)

	})
	t.Run("error after update data", func(t *testing.T) {
		repo.On("CekStok", mock.Anything, mock.Anything).Return(true).Once()
		repo.On("CreateUserInventoryData", mock.Anything, mock.Anything, mock.Anything).Return(input2).Once()
		repo.On("CreateUserDetailInventoryData", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(input).Once()
		repo.On("RekapStock", mock.Anything, mock.Anything, mock.Anything).Return(true).Once()
		repo.On("DeleteInOutBound", mock.Anything).Return("").Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateUserInventory(input2, 1)

		assert.Equal(t, 201, status)
		assert.Equal(t, input2, bound)

	})
}

func TestReadUserOutBoundDetail(t *testing.T) {
	repo := new(mocks.InventoryData)
	input := []domain.InventoryProduct{{ProductID: 2, Qty: 3, Unit: "kg"}, {ProductID: 5, Qty: 7, Unit: "kg"}}
	input3 := []domain.InventoryProduct{}

	t.Run("data not found", func(t *testing.T) {
		repo.On("ReadUserOutBoundDetailData", mock.Anything, mock.Anything).Return(input3).Once()
		useCase := New(repo, validator.New())
		bound, status, gen := useCase.ReadUserOutBoundDetail(1, "1345767564")

		assert.Equal(t, 200, status)
		assert.Equal(t, []domain.InventoryProduct{}, bound)
		assert.Equal(t, "", gen)

	})
	t.Run("succes", func(t *testing.T) {
		repo.On("ReadUserOutBoundDetailData", mock.Anything, mock.Anything).Return(input).Once()
		useCase := New(repo, validator.New())
		bound, status, gen := useCase.ReadUserOutBoundDetail(1, "1345767564")

		assert.Equal(t, 200, status)
		assert.Equal(t, input, bound)
		assert.Equal(t, "1345767564", gen)

	})
}

func TestReadUserOutBoundHistory(t *testing.T) {
	repo := new(mocks.InventoryData)
	input2 := []domain.Inventory{{ID: 2, OutBound: "6981234678", UserID: 5, InventoryProduct: []domain.InventoryProduct{}}}
	input4 := []domain.Inventory{}

	t.Run("data not found", func(t *testing.T) {
		repo.On("ReadUserOutBoundHistoryData", mock.Anything).Return(input4).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.ReadUserOutBoundHistory(1)

		assert.Equal(t, 200, status)
		assert.Equal(t, []domain.Inventory{}, bound)
	})
	t.Run("succes", func(t *testing.T) {
		repo.On("ReadUserOutBoundHistoryData", mock.Anything).Return(input2).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.ReadUserOutBoundHistory(1)

		assert.Equal(t, 200, status)
		assert.Equal(t, input2, bound)
	})
}

func TestAdminInventory(t *testing.T) {
	repo := new(mocks.InventoryData)
	input := []domain.InventoryProduct{{ProductID: 2, Qty: 3, Unit: "kg"}, {ProductID: 5, Qty: 7, Unit: "kg"}}
	input2 := domain.Inventory{ID: 2, OutBound: "6981234678", UserID: 5, InventoryProduct: []domain.InventoryProduct{}}
	input3 := []domain.InventoryProduct{}
	input4 := domain.Inventory{ID: 0, InventoryProduct: []domain.InventoryProduct{}}

	t.Run("error after creating outbound data", func(t *testing.T) {
		repo.On("CreateAdminInventoryData", mock.Anything, mock.Anything, mock.Anything).Return(input4).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateAdminInventory(input2, 1, "admin")

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.Inventory{}, bound)

	})

	t.Run("error after creating data", func(t *testing.T) {
		repo.On("CreateAdminInventoryData", mock.Anything, mock.Anything, mock.Anything).Return(input2).Once()
		repo.On("CreateAdminDetailInventoryData", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(input3).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateAdminInventory(input2, 1, "admin")

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.Inventory{}, bound)

	})

	t.Run("error after update data", func(t *testing.T) {
		repo.On("CreateAdminInventoryData", mock.Anything, mock.Anything, mock.Anything).Return(input2).Once()
		repo.On("CreateAdminDetailInventoryData", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(input).Once()
		repo.On("RekapAdminStock", mock.Anything, mock.Anything, mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateAdminInventory(input2, 1, "admin")

		assert.Equal(t, 404, status)
		assert.Equal(t, domain.Inventory{}, bound)

	})

	t.Run("error after update data", func(t *testing.T) {
		repo.On("CreateAdminInventoryData", mock.Anything, mock.Anything, mock.Anything).Return(input2).Once()
		repo.On("CreateAdminDetailInventoryData", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(input).Once()
		repo.On("RekapAdminStock", mock.Anything, mock.Anything, mock.Anything).Return(true).Once()
		repo.On("DeleteAdminInOutBound").Return("cannot delete data").Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateAdminInventory(input2, 1, "admin")

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.Inventory{}, bound)

	})
	t.Run("error after update data", func(t *testing.T) {
		repo.On("CreateAdminInventoryData", mock.Anything, mock.Anything, mock.Anything).Return(input2).Once()
		repo.On("CreateAdminDetailInventoryData", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(input).Once()
		repo.On("RekapAdminStock", mock.Anything, mock.Anything, mock.Anything).Return(true).Once()
		repo.On("DeleteAdminInOutBound").Return("no data deleted").Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateAdminInventory(input2, 1, "admin")

		assert.Equal(t, 404, status)
		assert.Equal(t, domain.Inventory{}, bound)

	})
	t.Run("error after update data", func(t *testing.T) {
		repo.On("CreateAdminInventoryData", mock.Anything, mock.Anything, mock.Anything).Return(input2).Once()
		repo.On("CreateAdminDetailInventoryData", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(input).Once()
		repo.On("RekapAdminStock", mock.Anything, mock.Anything, mock.Anything).Return(true).Once()
		repo.On("DeleteAdminInOutBound").Return("").Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.CreateAdminInventory(input2, 1, "admin")

		assert.Equal(t, 201, status)
		assert.Equal(t, input2, bound)

	})
}

func TestReadAdminOutBoundDetail(t *testing.T) {
	repo := new(mocks.InventoryData)
	input := []domain.InventoryProduct{{ProductID: 2, Qty: 3, Unit: "kg"}, {ProductID: 5, Qty: 7, Unit: "kg"}}
	input3 := []domain.InventoryProduct{}

	t.Run("data not found", func(t *testing.T) {
		repo.On("ReadAdminOutBoundDetailData", mock.Anything).Return(input3).Once()
		useCase := New(repo, validator.New())
		bound, status, gen := useCase.ReadAdminOutBoundDetail("123456789")

		assert.Equal(t, 200, status)
		assert.Equal(t, []domain.InventoryProduct{}, bound)
		assert.Equal(t, "", gen)

	})
	t.Run("succes", func(t *testing.T) {
		repo.On("ReadAdminOutBoundDetailData", mock.Anything).Return(input).Once()
		useCase := New(repo, validator.New())
		bound, status, gen := useCase.ReadAdminOutBoundDetail("123456789")

		assert.Equal(t, 200, status)
		assert.Equal(t, input, bound)
		assert.Equal(t, "123456789", gen)

	})
}

func TestReadAdminOutBoundHistory(t *testing.T) {
	repo := new(mocks.InventoryData)
	input2 := []domain.Inventory{{ID: 2, OutBound: "6981234678", UserID: 5, InventoryProduct: []domain.InventoryProduct{}}}
	input4 := []domain.Inventory{}

	t.Run("data not found", func(t *testing.T) {
		repo.On("ReadAdminOutBoundHistoryData").Return(input4).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.ReadAdminOutBoundHistory()

		assert.Equal(t, 200, status)
		assert.Equal(t, []domain.Inventory{}, bound)
	})
	t.Run("succes", func(t *testing.T) {
		repo.On("ReadAdminOutBoundHistoryData").Return(input2).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.ReadAdminOutBoundHistory()

		assert.Equal(t, 200, status)
		assert.Equal(t, input2, bound)
	})
}
