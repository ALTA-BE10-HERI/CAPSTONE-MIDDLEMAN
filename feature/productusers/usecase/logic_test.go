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

func TestCreateProduct(t *testing.T) {
	repo := new(mocks.ProductUserData)
	mockData := domain.ProductUser{ID: 1, IdUser: 2, Name: "tango", Unit: "piece", Stock: 100, Price: 3000, Image: "tango.jpg"}
	emptyMockData := domain.ProductUser{ID: 0, IdUser: 0, Name: "", Unit: "", Stock: 0, Price: 0, Image: ""}
	t.Run("Validasi", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res, status := useCase.CreateProduct(emptyMockData, 2)

		assert.Equal(t, 400, status)
		assert.Equal(t, domain.ProductUser{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("error after creating data", func(t *testing.T) {
		repo.On("CreateProductData", mock.Anything).Return(emptyMockData).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.CreateProduct(mockData, 2)

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.ProductUser{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("Succes", func(t *testing.T) {
		repo.On("CreateProductData", mock.Anything).Return(mockData).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.CreateProduct(mockData, 2)

		assert.Equal(t, 201, status)
		assert.Equal(t, mockData, res)
		repo.AssertExpectations(t)
	})
}

func TestReadAllProduct(t *testing.T) {
	repo := new(mocks.ProductUserData)
	outdata := []domain.ProductUser{{ID: 1, IdUser: 2, Name: "adidas", Unit: "kg", Stock: 10, Price: 10000, Image: "jpg"}}
	t.Run("data not found", func(t *testing.T) {
		repo.On("ReadAllProductData", mock.Anything).Return([]domain.ProductUser{}).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.ReadAllProduct(1)

		assert.Equal(t, 200, status)
		assert.Equal(t, []domain.ProductUser{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("succes", func(t *testing.T) {
		repo.On("ReadAllProductData", mock.Anything).Return(outdata).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.ReadAllProduct(1)

		assert.Equal(t, 200, status)
		assert.Equal(t, outdata, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateProduct(t *testing.T) {
	repo := new(mocks.ProductUserData)
	insert := domain.ProductUser{ID: 9, IdUser: 10, Name: "adidas", Unit: "kg", Stock: 10, Price: 10000, Image: "ini gambar"}
	insert2 := domain.ProductUser{Name: "", Unit: "", Stock: 0, Price: 0, Image: ""}
	insert3 := domain.ProductUser{ID: 0, IdUser: 0, Name: "adidas", Unit: "kg", Stock: 10, Price: 10000, Image: "ini gambar"}

	t.Run("Filled incorrectly", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res, err := useCase.UpdateProduct(insert2, 1, 1)

		assert.Equal(t, 400, err)
		assert.Equal(t, domain.ProductUser{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("Empty Data", func(t *testing.T) {
		repo.On("UpdateProductData", mock.Anything, mock.Anything, mock.Anything).Return(insert3).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.UpdateProduct(insert3, 1, 1)

		assert.Equal(t, 404, err)
		assert.Equal(t, domain.ProductUser{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("Succes", func(t *testing.T) {
		repo.On("UpdateProductData", mock.Anything, mock.Anything, mock.Anything).Return(insert).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.UpdateProduct(insert3, 1, 1)

		assert.Equal(t, 200, err)
		assert.Equal(t, insert, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteProduct(t *testing.T) {
	repo := new(mocks.ProductUserData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("DeleteProductData", mock.Anything, mock.Anything).Return("").Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeleteProduct(1, 1)

		assert.Equal(t, 204, delete)
		repo.AssertExpectations(t)
	})

	t.Run("cannot delete data", func(t *testing.T) {
		repo.On("DeleteProductData", mock.Anything, mock.Anything).Return("cannot delete data").Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeleteProduct(1, 1)

		assert.Equal(t, 500, delete)
		repo.AssertExpectations(t)
	})

	t.Run("no data deleted", func(t *testing.T) {
		repo.On("DeleteProductData", mock.Anything, mock.Anything).Return("no data deleted").Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeleteProduct(1, 1)

		assert.Equal(t, 404, delete)
		repo.AssertExpectations(t)
	})
}

func TestSearchRestoBusiness(t *testing.T) {
	repo := new(mocks.ProductUserData)
	outdata := []domain.ProductUser{{ID: 1, IdUser: 2, Name: "Beras", Unit: "kg", Stock: 10, Price: 10000, Image: "jpg"}}
	t.Run("Succes delete", func(t *testing.T) {
		repo.On("SearchRestoData", "Beras").Return(outdata, nil).Once()
		usecase := New(repo, validator.New())
		res, err := usecase.SearchRestoBusiness("Beras")

		assert.NoError(t, err)
		assert.Equal(t, outdata, res)
		repo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		repo.On("SearchRestoData", "Beras").Return([]domain.ProductUser{}, errors.New("error")).Once()
		usecase := New(repo, validator.New())
		res, err := usecase.SearchRestoBusiness("Beras")

		assert.Error(t, err)
		assert.Equal(t, []domain.ProductUser{}, res)
		repo.AssertExpectations(t)
	})
}
