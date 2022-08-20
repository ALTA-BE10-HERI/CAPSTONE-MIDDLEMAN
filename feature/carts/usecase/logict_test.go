package usecase

import (
	"errors"
	"middleman-capstone/domain"
	"middleman-capstone/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateData(t *testing.T) {
	repo := new(mocks.ChartData)
	productStock := 2000
	productPrice := 2000
	cekCurrentQty := 1000
	mockData := domain.Cart{ID: 1, Qty: 1000, Subtotal: 10900, Grandtotal: 20000, UserID: 2, ProductID: 9}
	mockData2 := domain.Cart{ID: 0, Qty: 0, Subtotal: 0, Grandtotal: 0, UserID: 0, ProductID: 0}

	t.Run("All data not filled", func(t *testing.T) {
		useCase := New(repo)
		res, err := useCase.CreateData(mockData2)
		assert.Equal(t, -1, res)
		assert.Error(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Insufficient Stock", func(t *testing.T) {
		repo.On("GetStockProduct", mock.Anything).Return(productStock).Once()
		useCase := New(repo)
		res, err := useCase.CreateData(mockData)
		if mockData.Qty > productStock {
			assert.Equal(t, 400, res)
			assert.Error(t, err)
			repo.AssertExpectations(t)
		}
	})

	t.Run("cek cart true", func(t *testing.T) {
		repo.On("CheckCart", mock.Anything, mock.Anything).Return(true, 1, 1).Once()
		repo.On("GetQtyProductCart", 1).Return(cekCurrentQty).Once()
		useCase := New(repo)
		res, err := useCase.CreateData(mockData)

		if (cekCurrentQty + mockData.Qty) > productStock {
			assert.Equal(t, 400, res)
			assert.Error(t, err)
			repo.AssertExpectations(t)
		}

	})

	t.Run("succes", func(t *testing.T) {
		useCase := New(repo)
		res, err := useCase.CreateData(mockData)
		if mockData.Qty == 9 || mockData.ProductID == 10 {

			repo.On("GetStockProduct", mock.Anything).Return(productStock).Once()
			repo.On("GetPriceProduct", mock.Anything).Return(productPrice).Once()
			repo.On("CheckCart", mock.Anything, mock.Anything).Return(true, 1, 1).Once()
			repo.On("GetQtyProductCart", 1).Return(cekCurrentQty).Once()
			repo.On("UpdateDataDB", mock.Anything, mock.Anything, mock.Anything).Return(1, nil).Once()

			assert.Equal(t, 1, res)
			assert.NoError(t, err)
			repo.AssertExpectations(t)
		}

	})

}

func TestGetAllData(t *testing.T) {
	repo := new(mocks.ChartData)
	limit := 1
	offset := 10
	id := 8
	outdata := []domain.Cart{{ID: 1, Qty: 1000, Subtotal: 10900, Grandtotal: 20000, UserID: 2, ProductID: 9}}
	t.Run("Success Get product", func(t *testing.T) {
		repo.On("SelectData", limit, offset, id).Return(outdata, nil).Once()
		useCase := New(repo)
		res, total, err := useCase.GetAllData(limit, offset, id)

		assert.Equal(t, outdata, res)
		assert.Equal(t, 10900, total)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("SelectData", limit, offset, id).Return([]domain.Cart{}, errors.New("error")).Once()
		useCase := New(repo)
		res, total, err := useCase.GetAllData(limit, offset, id)

		assert.NotNil(t, err)
		assert.Equal(t, []domain.Cart{}, res)
		assert.Equal(t, 0, total)
		repo.AssertExpectations(t)
	})
}

func TestUpdateData(t *testing.T) {
	// repo := new(mocks.ChartData)
	// insert := domain.Cart{Qty: 1000}

	// t.Run("Success Update product", func(t *testing.T) {
	// 	repo.On("GetPriceProduct", mock.Anything).Return(1000).Once()
	// 	repo.On("UpdateDataDB", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(1, nil).Once()
	// 	useCase := New(repo)
	// 	res, err := useCase.UpdateData(1000, 1, 1)

	// 	assert.NoError(t, err)
	// 	assert.Equal(t, 1, res)
	// 	repo.AssertExpectations(t)
	// })

	// t.Run("Failed Update product", func(t *testing.T) {
	// 	repo.On("GetPriceProduct", mock.Anything).Return(insert.Qty).Once()
	// 	repo.On("UpdateDataDB", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(0, errors.New("error")).Once()
	// 	useCase := New(repo)
	// 	res, err := useCase.UpdateData(0, 1, 1)

	// 	assert.Error(t, err)
	// 	assert.Equal(t, 0, res)
	// 	repo.AssertExpectations(t)
	// })

}

func TestDeleteData(t *testing.T) {
	repo := new(mocks.ChartData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("DeleteDataDB", mock.Anything, mock.Anything).Return(1, nil).Once()
		usecase := New(repo)
		res, err := usecase.DeleteData(1, 1)

		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("No Data Found", func(t *testing.T) {
		repo.On("DeleteDataDB", mock.Anything, mock.Anything).Return(0, errors.New("error")).Once()
		usecase := New(repo)
		res, err := usecase.DeleteData(0, 0)

		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})

}
