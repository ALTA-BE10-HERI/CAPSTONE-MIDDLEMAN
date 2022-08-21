package usecase

import (
	"errors"
	"middleman-capstone/domain"
	"middleman-capstone/domain/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllAdmin(t *testing.T) {
	repo := new(mocks.OrderData)
	input := []domain.Order{{ID: 1, UserID: 1, GrandTotal: 100000, Status: "waiting confirmation", PaymentLink: "dimana", OrderName: "123456", Items: []domain.Items{}}}
	input3 := []domain.Order{}

	t.Run("not admin", func(t *testing.T) {
		useCase := New(repo, validator.New())
		order, err := useCase.GetAllAdmin(1, 1, "user")

		assert.NoError(t, err)
		assert.Equal(t, []domain.Order{}, order)

	})

	t.Run("select admin failed", func(t *testing.T) {
		repo.On("SelectDataAdminAll", mock.Anything, mock.Anything).Return(input3, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		order, err := useCase.GetAllAdmin(1, 1, "admin")

		assert.Error(t, err)
		assert.Equal(t, []domain.Order{}, order)

	})

	t.Run("select admin success", func(t *testing.T) {
		repo.On("SelectDataAdminAll", mock.Anything, mock.Anything).Return(input, nil).Once()
		useCase := New(repo, validator.New())
		order, err := useCase.GetAllAdmin(1, 1, "admin")

		assert.NoError(t, err)
		assert.Equal(t, input, order)

	})
}

func TestGetAllUser(t *testing.T) {
	repo := new(mocks.OrderData)
	input := []domain.Order{{ID: 1, UserID: 1, GrandTotal: 100000, Status: "waiting confirmation", PaymentLink: "dimana", OrderName: "123456", Items: []domain.Items{}}}
	input3 := []domain.Order{}

	t.Run("id user = 0", func(t *testing.T) {
		useCase := New(repo, validator.New())
		order, err := useCase.GetAllUser(1, 1, 0)

		assert.NoError(t, err)
		assert.Equal(t, input3, order)
	})

	t.Run("failed select data user all", func(t *testing.T) {
		repo.On("SelectDataUserAll", mock.Anything, mock.Anything, mock.Anything).Return(input3, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		order, err := useCase.GetAllUser(1, 1, 1)

		assert.NoError(t, err)
		assert.Equal(t, input3, order)
	})

	t.Run("succes select data user all", func(t *testing.T) {
		repo.On("SelectDataUserAll", mock.Anything, mock.Anything, mock.Anything).Return(input, nil).Once()
		useCase := New(repo, validator.New())
		order, err := useCase.GetAllUser(1, 1, 1)

		assert.NoError(t, err)
		assert.Equal(t, input, order)
	})
}

func TestGetDetail(t *testing.T) {
	repo := new(mocks.OrderData)

	t.Run("grandtotal = 0", func(t *testing.T) {
		repo.On("GetDetailData", mock.Anything, mock.Anything).Return(0, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		order, err := useCase.GetDetail(1, 1)

		assert.NoError(t, err)
		assert.Equal(t, -1, order)
	})

	t.Run("grandtotal != 0, error != nil", func(t *testing.T) {
		repo.On("GetDetailData", mock.Anything, mock.Anything).Return(10000, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		order, err := useCase.GetDetail(1, 1)

		assert.NoError(t, err)
		assert.Equal(t, 400, order)
	})

	t.Run("succes", func(t *testing.T) {
		repo.On("GetDetailData", mock.Anything, mock.Anything).Return(10000, nil).Once()
		useCase := New(repo, validator.New())
		order, err := useCase.GetDetail(1, 1)

		assert.NoError(t, err)
		assert.Equal(t, 10000, order)
	})
}

func TestGetItems(t *testing.T) {
	repo := new(mocks.OrderData)
	input := []domain.Items{{ID: 1, OrderID: 1, ProductID: 1, ProductName: "beras", Subtotal: 50000, Unit: "kg", Qty: 1}}
	input3 := []domain.Items{}

	t.Run("failed to get detail items", func(t *testing.T) {
		repo.On("GetDetailItems", mock.Anything).Return(input3, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		order, err := useCase.GetItems(1)

		assert.NoError(t, err)
		assert.Equal(t, input3, order)
	})

	t.Run("succes to get detail items", func(t *testing.T) {
		repo.On("GetDetailItems", mock.Anything).Return(input, nil).Once()
		useCase := New(repo, validator.New())
		order, err := useCase.GetItems(1)

		assert.NoError(t, err)
		assert.Equal(t, input, order)
	})
}

func TestGetIncoming(t *testing.T) {
	repo := new(mocks.OrderData)
	input := []domain.Order{{ID: 1, UserID: 1, GrandTotal: 100000, Status: "waiting confirmation", PaymentLink: "dimana", OrderName: "123456", Items: []domain.Items{}}}
	input3 := []domain.Order{}

	t.Run("role not admin", func(t *testing.T) {
		useCase := New(repo, validator.New())
		order, err := useCase.GetIncoming(1, 1, "user")

		assert.Error(t, err)
		assert.Equal(t, input3, order)
	})

	t.Run("failed get incoming data", func(t *testing.T) {
		repo.On("GetIncomingData", mock.Anything, mock.Anything).Return(input3, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		order, err := useCase.GetIncoming(1, 1, "admin")

		assert.NoError(t, err)
		assert.Equal(t, input3, order)
	})

	t.Run("succes get incoming data", func(t *testing.T) {
		repo.On("GetIncomingData", mock.Anything, mock.Anything).Return(input, nil).Once()
		useCase := New(repo, validator.New())
		order, err := useCase.GetIncoming(1, 1, "admin")

		assert.NoError(t, err)
		assert.Equal(t, input, order)
	})

}

func TestCreateOrder(t *testing.T) {
	repo := new(mocks.OrderData)
	input := []domain.Items{{ID: 1, OrderID: 1, ProductID: 1, ProductName: "beras", Subtotal: 50000, Unit: "kg", Qty: 1}}
	input3 := []domain.Items{}
	input4 := domain.Order{ID: 0, UserID: 0, GrandTotal: 0, Status: "", PaymentLink: "", OrderName: "", Items: []domain.Items{{ID: 0, OrderID: 0, ProductID: 0, ProductName: "", Subtotal: 0, Unit: "", Qty: 0}}}

	t.Run("failed to insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(1, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		status := useCase.CreateOrder(input4, 1)

		assert.Equal(t, 400, status)
	})

	t.Run("failed to create", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(1, nil).Once()
		repo.On("InsertData", mock.Anything, mock.Anything).Return(input3).Once()
		useCase := New(repo, validator.New())
		status := useCase.CreateOrder(input4, 1)

		assert.Equal(t, 500, status)
	})

	t.Run("succes to create", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(1, nil).Once()
		repo.On("InsertData", mock.Anything, mock.Anything).Return(input).Once()
		useCase := New(repo, validator.New())
		status := useCase.CreateOrder(input4, 1)

		assert.Equal(t, 201, status)
	})

}

func TestPayment(t *testing.T) {
	repo := new(mocks.OrderData)
	return1 := domain.User{ID: 0}
	return2 := domain.User{ID: 1}

	t.Run("failed to get user", func(t *testing.T) {
		repo.On("GetUser", mock.Anything).Return(return1, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		name, link, token, user := useCase.Payment(10000, 0)

		assert.Equal(t, "", name)
		assert.Equal(t, "", link)
		assert.Equal(t, "", token)
		assert.Equal(t, return1, user)
	})

	t.Run("succes to get user", func(t *testing.T) {
		repo.On("GetUser", mock.Anything).Return(return2, nil).Once()
		useCase := New(repo, validator.New())
		_, link, token, user := useCase.Payment(10000, 1)

		assert.Equal(t, "", link)
		assert.Equal(t, "", token)
		assert.Equal(t, return2, user)
	})
}

func TestAcceptPayment(t *testing.T) {
	repo := new(mocks.OrderData)
	return1 := domain.PaymentWeb{TransactionStatus: "settlement", OrderName: "123"}
	return2 := domain.PaymentWeb{TransactionStatus: "expire", OrderName: "123"}
	return3 := domain.PaymentWeb{TransactionStatus: "cancel", OrderName: "123"}
	return4 := domain.PaymentWeb{TransactionStatus: ""}

	t.Run("data kosong", func(t *testing.T) {
		useCase := New(repo, validator.New())
		row, err := useCase.AcceptPayment(return4)

		assert.Equal(t, -1, row)
		assert.NoError(t, err)
	})

	t.Run("data settlement", func(t *testing.T) {
		repo.On("AcceptPaymentData", mock.Anything).Return(1, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		row, err := useCase.AcceptPayment(return1)

		assert.Error(t, err)
		assert.Equal(t, 500, row)
		repo.AssertExpectations(t)
	})

	t.Run("data settlement", func(t *testing.T) {
		repo.On("AcceptPaymentData", mock.Anything).Return(1, nil).Once()
		useCase := New(repo, validator.New())
		_, err := useCase.AcceptPayment(return1)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("data expire", func(t *testing.T) {
		repo.On("CancelPaymentData", mock.Anything).Return(1, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		row, err := useCase.AcceptPayment(return2)

		assert.Error(t, err)
		assert.Equal(t, 500, row)
		repo.AssertExpectations(t)
	})

	t.Run("data expire", func(t *testing.T) {
		repo.On("CancelPaymentData", mock.Anything).Return(1, nil).Once()
		useCase := New(repo, validator.New())
		_, err := useCase.AcceptPayment(return2)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("data cancel", func(t *testing.T) {
		repo.On("CancelPaymentData", mock.Anything).Return(1, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		row, err := useCase.AcceptPayment(return3)

		assert.Error(t, err)
		assert.Equal(t, 500, row)
		repo.AssertExpectations(t)
	})

	t.Run("data cancel", func(t *testing.T) {
		repo.On("CancelPaymentData", mock.Anything).Return(1, nil).Once()
		useCase := New(repo, validator.New())
		_, err := useCase.AcceptPayment(return3)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

}

func TestConfirmOrder(t *testing.T) {
	repo := new(mocks.OrderData)
	return1 := domain.Order{}
	return2 := domain.Order{ID: 0}
	return3 := domain.Order{ID: 1, UserID: 2, GrandTotal: 100000, Status: "delivered", PaymentLink: "jpg", OrderName: "123", Items: []domain.Items{}}
	return4 := domain.User{}
	return5 := domain.User{ID: 1}

	t.Run("you dont have access", func(t *testing.T) {
		useCase := New(repo, validator.New())
		order, status := useCase.ConfirmOrder("123", "user")

		assert.Equal(t, return1, order)
		assert.Equal(t, 401, status)
	})

	t.Run("failed update status", func(t *testing.T) {
		repo.On("ConfirmOrderData", mock.Anything).Return(return2).Once()
		useCase := New(repo, validator.New())
		order, status := useCase.ConfirmOrder("123", "admin")

		assert.Equal(t, return1, order)
		assert.Equal(t, 404, status)
	})

	t.Run("failed get user", func(t *testing.T) {
		repo.On("ConfirmOrderData", mock.Anything).Return(return3).Once()
		repo.On("GetUser", mock.Anything).Return(return4, nil).Once()
		useCase := New(repo, validator.New())
		order, status := useCase.ConfirmOrder("123", "admin")

		assert.Equal(t, return1, order)
		assert.Equal(t, 500, status)
	})
	t.Run("failed get user", func(t *testing.T) {
		repo.On("ConfirmOrderData", mock.Anything).Return(return3).Once()
		repo.On("GetUser", mock.Anything).Return(return5, nil).Once()
		useCase := New(repo, validator.New())
		order, status := useCase.ConfirmOrder("123", "admin")

		assert.Equal(t, return3, order)
		assert.Equal(t, 200, status)
	})
}

func TestDoneOrder(t *testing.T) {
	repo := new(mocks.OrderData)
	return0 := domain.Order{}
	return1 := domain.Order{ID: 0}
	// return2 := domain.Order{ID: 1}
	return3 := []domain.Items{}
	return4 := []domain.Items{{ID: 2, OrderID: 2, ProductID: 2, ProductName: "Name", Subtotal: 12000, Unit: "kg", Qty: 12}}
	return5 := domain.Order{ID: 1, UserID: 2, GrandTotal: 100000, Status: "delivered", PaymentLink: "jpg", OrderName: "123", Items: []domain.Items{}}
	t.Run("failed to update data", func(t *testing.T) {
		repo.On("DoneOrderData", mock.Anything).Return(return1).Once()
		useCase := New(repo, validator.New())
		order, status := useCase.DoneOrder("123")

		assert.Equal(t, return0, order)
		assert.Equal(t, 404, status)
	})

	t.Run("failed to delete cart", func(t *testing.T) {
		repo.On("DoneOrderData", mock.Anything).Return(return5).Once()
		repo.On("DeleteCart", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		order, status := useCase.DoneOrder("123")

		assert.Equal(t, return0, order)
		assert.Equal(t, 500, status)
	})

	t.Run("failed to update admin stok", func(t *testing.T) {
		repo.On("DoneOrderData", mock.Anything).Return(return5).Once()
		repo.On("DeleteCart", mock.Anything).Return(true).Once()
		repo.On("UpdateStokAdmin", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		order, status := useCase.DoneOrder("123")

		assert.Equal(t, return0, order)
		assert.Equal(t, 500, status)
	})

	t.Run("failed cek user", func(t *testing.T) {
		repo.On("DoneOrderData", mock.Anything).Return(return5).Once()
		repo.On("DeleteCart", mock.Anything).Return(true).Once()
		repo.On("UpdateStokAdmin", mock.Anything).Return(true).Once()
		repo.On("CekUser", mock.Anything).Return(return3, 1).Once()
		useCase := New(repo, validator.New())
		order, status := useCase.DoneOrder("123")

		assert.Equal(t, return0, order)
		assert.Equal(t, 500, status)
	})

	t.Run("failed to cek own and create", func(t *testing.T) {
		repo.On("DoneOrderData", mock.Anything).Return(return5).Once()
		repo.On("DeleteCart", mock.Anything).Return(true).Once()
		repo.On("UpdateStokAdmin", mock.Anything).Return(true).Once()
		repo.On("CekUser", mock.Anything).Return(return4, 1).Once()
		repo.On("CekOwned", mock.Anything, mock.Anything).Return(false).Once()
		repo.On("CreateNewProduct", mock.Anything, mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		order, status := useCase.DoneOrder("123")

		assert.Equal(t, return0, order)
		assert.Equal(t, 500, status)
	})

	t.Run("succes to cek own and create", func(t *testing.T) {
		repo.On("DoneOrderData", mock.Anything).Return(return5).Once()
		repo.On("DeleteCart", mock.Anything).Return(true).Once()
		repo.On("UpdateStokAdmin", mock.Anything).Return(true).Once()
		repo.On("CekUser", mock.Anything).Return(return4, 1).Once()
		repo.On("CekOwned", mock.Anything, mock.Anything).Return(false).Once()
		repo.On("CreateNewProduct", mock.Anything, mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		order, status := useCase.DoneOrder("123")

		assert.Equal(t, return5, order)
		assert.Equal(t, 200, status)
	})

	t.Run("failed to cek own and update", func(t *testing.T) {
		repo.On("DoneOrderData", mock.Anything).Return(return5).Once()
		repo.On("DeleteCart", mock.Anything).Return(true).Once()
		repo.On("UpdateStokAdmin", mock.Anything).Return(true).Once()
		repo.On("CekUser", mock.Anything).Return(return4, 1).Once()
		repo.On("CekOwned", mock.Anything, mock.Anything).Return(true).Once()
		repo.On("UpdateNewProduct", mock.Anything, mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		order, status := useCase.DoneOrder("123")

		assert.Equal(t, return0, order)
		assert.Equal(t, 500, status)
	})

	t.Run("succes to cek own and update", func(t *testing.T) {
		repo.On("DoneOrderData", mock.Anything).Return(return5).Once()
		repo.On("DeleteCart", mock.Anything).Return(true).Once()
		repo.On("UpdateStokAdmin", mock.Anything).Return(true).Once()
		repo.On("CekUser", mock.Anything).Return(return4, 1).Once()
		repo.On("CekOwned", mock.Anything, mock.Anything).Return(true).Once()
		repo.On("UpdateNewProduct", mock.Anything, mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		order, status := useCase.DoneOrder("123")

		assert.Equal(t, return5, order)
		assert.Equal(t, 200, status)
	})

}
