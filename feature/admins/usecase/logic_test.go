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

func TestProduct(t *testing.T) {
	repo := new(mocks.ProductData)
	mockData := domain.Product{ID: 1, IdAdmin: 2, Name: "tango", Unit: "piece", Stock: 100, Price: 3000, Image: "tango.jpg"}
	emptyMockData := domain.Product{ID: 0, IdAdmin: 0, Name: "", Unit: "", Stock: 0, Price: 0, Image: ""}
	t.Run("Success insert product", func(t *testing.T) {
		repo.On("CreateProductData", mock.Anything).Return(mockData).Once()
		useCase := New(repo, validator.New())
		res := useCase.CreateProduct(mockData, 2)

		assert.Equal(t, 201, res)
		assert.Greater(t, mockData.ID, 0)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("CreateProductData", mock.Anything).Return(emptyMockData).Once()
		useCase := New(repo, validator.New())
		res := useCase.CreateProduct(emptyMockData, 5)

		assert.Equal(t, 400, res)
		assert.Equal(t, emptyMockData.ID, 0)
	})

	t.Run("Internal server error", func(t *testing.T) {
		mockData.ID = 0
		repo.On("CreateProductData", mock.Anything).Return(mockData).Once()
		useCase := New(repo, validator.New())
		res := useCase.CreateProduct(mockData, 5)

		assert.Equal(t, 500, res)
		assert.Equal(t, emptyMockData.ID, 0)
	})
}

func TestGetAllData(t *testing.T) {
	repo := new(mocks.ProductData)
	limit := 1
	offset := 10
	outdata := []domain.Product{{ID: 1, IdAdmin: 2, Name: "adidas", Unit: "kg", Stock: 10, Price: 10000, Image: "jpg"}}
	t.Run("Success Get product", func(t *testing.T) {
		repo.On("GetAllProductData", limit, offset).Return(outdata, nil).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.GetAllProduct(limit, offset)

		assert.NoError(t, err)
		assert.Equal(t, outdata, res)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("GetAllProductData", limit, offset).Return([]domain.Product{}, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.GetAllProduct(limit, offset)

		assert.NotNil(t, err)
		assert.Equal(t, []domain.Product{}, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateProduct(t *testing.T) {
	repo := new(mocks.ProductData)
	insert := domain.Product{Name: "adidas", Unit: "kg", Stock: 10, Price: 10000, Image: "ini gambar"}

	t.Run("Success Update product", func(t *testing.T) {
		repo.On("UpdateProductData", mock.Anything, 1).Return(1, nil).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.UpdateProduct(insert, 1)

		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update product", func(t *testing.T) {
		repo.On("UpdateProductData", mock.Anything, 1).Return(0, errors.New("error")).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.UpdateProduct(insert, 1)

		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})

}

func TestDeleteProduct(t *testing.T) {
	repo := new(mocks.ProductData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("DeleteProductData", 1).Return(1, nil).Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeleteProduct(1)

		assert.Equal(t, 204, delete)
		repo.AssertExpectations(t)
	})

	t.Run("No Data Found", func(t *testing.T) {
		repo.On("DeleteProductData", 1).Return(1, errors.New("error")).Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeleteProduct(1)

		assert.Equal(t, 404, delete)
		repo.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		repo.On("DeleteProductData", 1).Return(0, nil).Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeleteProduct(1)

		assert.Equal(t, 500, delete)
		repo.AssertExpectations(t)
	})
}

func TestSearchRestoBusiness(t *testing.T) {
	repo := new(mocks.ProductData)
	outdata := []domain.Product{{ID: 1, IdAdmin: 2, Name: "Beras", Unit: "kg", Stock: 10, Price: 10000, Image: "jpg"}}
	t.Run("Succes delete", func(t *testing.T) {
		repo.On("SearchRestoData", "Beras").Return(outdata, nil).Once()
		usecase := New(repo, validator.New())
		res, err := usecase.SearchRestoBusiness("Beras")

		assert.NoError(t, err)
		assert.Equal(t, outdata, res)
		repo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		repo.On("SearchRestoData", "Beras").Return([]domain.Product{}, errors.New("error")).Once()
		usecase := New(repo, validator.New())
		res, err := usecase.SearchRestoBusiness("Beras")

		assert.Error(t, err)
		assert.Equal(t, []domain.Product{}, res)
		repo.AssertExpectations(t)
	})
}
