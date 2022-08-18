package usecase

import (
	"middleman-capstone/domain"
	"middleman-capstone/domain/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateData(t *testing.T) {
	repo := new(mocks.InOutBoundData)
	input := domain.InOutBounds{IdProduct: 5, Qty: 1}
	returndata := domain.InOutBounds{ID: 1, IdProduct: 5, Name: "tango", Unit: "kg", Qty: 9}
	// mockData := domain.InOutBounds{ID: 1, IdUser: 2, IdProduct: 5, Name: "tango", Unit: "kg", Qty: 1000, Role: "admin"}
	emptymockData := domain.InOutBounds{ID: 0, IdUser: 0, IdProduct: 0, Name: "", Unit: "", Qty: 0}

	t.Run("Data Not Found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res, status := useCase.AddEntry(emptymockData, 0, "")

		assert.Equal(t, 400, status)
		assert.Equal(t, domain.InOutBounds{}, res)
	})

	t.Run("Wrong by Admin 1.1", func(t *testing.T) {
		repo.On("CekOwnerAdminEntry", mock.Anything).Return(true).Once()
		repo.On("CekAdminEntry", mock.Anything).Return(true, 1, 20).Once()
		repo.On("UpdateQty", mock.Anything, mock.Anything).Return(emptymockData).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.AddEntry(input, 1, "admin")

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.InOutBounds{}, bound)

	})

	t.Run("Wrong by Admin 1.2", func(t *testing.T) {
		repo.On("CekOwnerAdminEntry", mock.Anything).Return(true).Once()
		repo.On("CekAdminEntry", mock.Anything).Return(true, 1, 20).Once()
		repo.On("UpdateQty", mock.Anything, mock.Anything).Return(returndata).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.AddEntry(input, 1, "admin")

		assert.Equal(t, 201, status)
		assert.Equal(t, returndata, bound)

	})

	t.Run("Wrong by Admin 2.1", func(t *testing.T) {
		repo.On("CekOwnerAdminEntry", mock.Anything).Return(true).Once()
		repo.On("CekAdminEntry", mock.Anything).Return(false, 1, 20).Once()
		repo.On("AddEntryData", mock.Anything).Return(emptymockData).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.AddEntry(input, 1, "admin")

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.InOutBounds{}, bound)

	})

	t.Run("Wrong by Admin 2.2", func(t *testing.T) {
		repo.On("CekOwnerAdminEntry", mock.Anything).Return(true).Once()
		repo.On("CekAdminEntry", mock.Anything).Return(false, 1, 20).Once()
		repo.On("AddEntryData", mock.Anything).Return(returndata).Once()
		repo.On("UpdateEntryAdminData", mock.Anything).Return(emptymockData).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.AddEntry(input, 1, "admin")

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.InOutBounds{}, bound)

	})

	t.Run("Wrong by Admin 2.3", func(t *testing.T) {
		repo.On("CekOwnerAdminEntry", mock.Anything).Return(true).Once()
		repo.On("CekAdminEntry", mock.Anything).Return(false, 1, 20).Once()
		repo.On("AddEntryData", mock.Anything).Return(returndata).Once()
		repo.On("UpdateEntryAdminData", mock.Anything).Return(returndata).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.AddEntry(input, 1, "admin")

		assert.Equal(t, 201, status)
		assert.Equal(t, returndata, bound)

	})

	t.Run("Wrong by Admin 3.1", func(t *testing.T) {
		repo.On("CekOwnerAdminEntry", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.AddEntry(input, 1, "admin")

		assert.Equal(t, 404, status)
		assert.Equal(t, domain.InOutBounds{}, bound)

	})

	t.Run("Wrong by User 1.1", func(t *testing.T) {
		repo.On("CekOwnerEntry", mock.Anything).Return(true).Once()
		repo.On("CekUserEntry", mock.Anything).Return(true, 1, 20).Once()
		repo.On("UpdateQty", mock.Anything, mock.Anything).Return(emptymockData).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.AddEntry(input, 1, "user")

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.InOutBounds{}, bound)

	})

	t.Run("Wrong by User 1.2", func(t *testing.T) {
		repo.On("CekOwnerEntry", mock.Anything).Return(true).Once()
		repo.On("CekUserEntry", mock.Anything).Return(true, 1, 20).Once()
		repo.On("UpdateQty", mock.Anything, mock.Anything).Return(returndata).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.AddEntry(input, 1, "user")

		assert.Equal(t, 201, status)
		assert.Equal(t, returndata, bound)

	})

	t.Run("Wrong by User 2.1", func(t *testing.T) {
		repo.On("CekOwnerEntry", mock.Anything).Return(true).Once()
		repo.On("CekUserEntry", mock.Anything).Return(false, 1, 20).Once()
		repo.On("AddEntryData", mock.Anything).Return(emptymockData).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.AddEntry(input, 1, "user")

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.InOutBounds{}, bound)

	})

	t.Run("Wrong by User 2.2", func(t *testing.T) {
		repo.On("CekOwnerEntry", mock.Anything).Return(true).Once()
		repo.On("CekUserEntry", mock.Anything).Return(false, 1, 20).Once()
		repo.On("AddEntryData", mock.Anything).Return(returndata).Once()
		repo.On("UpdateEntryUserData", mock.Anything, mock.Anything).Return(emptymockData).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.AddEntry(input, 1, "user")

		assert.Equal(t, 500, status)
		assert.Equal(t, domain.InOutBounds{}, bound)

	})

	t.Run("Wrong by User 2.3", func(t *testing.T) {
		repo.On("CekOwnerEntry", mock.Anything).Return(true).Once()
		repo.On("CekUserEntry", mock.Anything).Return(false, 1, 20).Once()
		repo.On("AddEntryData", mock.Anything).Return(returndata).Once()
		repo.On("UpdateEntryUserData", mock.Anything, mock.Anything).Return(returndata).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.AddEntry(input, 1, "user")

		assert.Equal(t, 201, status)
		assert.Equal(t, returndata, bound)

	})

	t.Run("Wrong by User 3.1", func(t *testing.T) {
		repo.On("CekOwnerEntry", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.AddEntry(input, 1, "user")

		assert.Equal(t, 404, status)
		assert.Equal(t, domain.InOutBounds{}, bound)

	})
}

func TestReadEntry(t *testing.T) {
	repo := new(mocks.InOutBoundData)
	returndata := []domain.InOutBounds{{ID: 1, IdProduct: 5, Name: "tango", Unit: "kg", Qty: 9}}
	returndata2 := []domain.InOutBounds{}

	t.Run("Wrong by Admin 1.1", func(t *testing.T) {
		repo.On("ReadEntryAdminData", mock.Anything).Return(returndata2).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.ReadEntry(1, "admin")

		assert.Equal(t, 200, status)
		assert.Equal(t, []domain.InOutBounds{}, bound)

	})
	t.Run("Wrong by Admin 1.2", func(t *testing.T) {
		repo.On("ReadEntryAdminData", mock.Anything).Return(returndata).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.ReadEntry(1, "admin")

		assert.Equal(t, 200, status)
		assert.Equal(t, returndata, bound)

	})

	t.Run("Wrong by User 1.1", func(t *testing.T) {
		repo.On("ReadEntryUserData", mock.Anything).Return(returndata2).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.ReadEntry(1, "user")

		assert.Equal(t, 200, status)
		assert.Equal(t, []domain.InOutBounds{}, bound)

	})
	t.Run("Wrong by Admin 1.2", func(t *testing.T) {
		repo.On("ReadEntryUserData", mock.Anything).Return(returndata).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.ReadEntry(1, "user")

		assert.Equal(t, 200, status)
		assert.Equal(t, returndata, bound)

	})
}

func TestUpdateEntry(t *testing.T) {
	repo := new(mocks.InOutBoundData)
	input := domain.InOutBounds{Qty: 1}
	// returndata := domain.InOutBounds{ID: 1, IdProduct: 5, Name: "tango", Unit: "kg", Qty: 9}
	mockData := domain.InOutBounds{ID: 1, IdUser: 2, IdProduct: 5, Name: "tango", Unit: "kg", Qty: 1000}
	mockData2 := domain.InOutBounds{ID: 1, IdUser: 2, IdProduct: 5, Name: "tango", Unit: "kg", Qty: 1000, Note: "insufficient stok"}
	emptymockData := domain.InOutBounds{ID: 0, IdUser: 0, IdProduct: 0, Name: "", Unit: "", Qty: 0}

	t.Run("Data Not Found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res, status := useCase.UpdateEntry(emptymockData, 0, 0, "")

		assert.Equal(t, 400, status)
		assert.Equal(t, domain.InOutBounds{}, res)
	})

	t.Run("Wrong by Admin 1.1", func(t *testing.T) {
		repo.On("UpdateQtyAdminData", mock.Anything).Return(emptymockData).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.UpdateEntry(input, 1, 1, "admin")

		assert.Equal(t, 404, status)
		assert.Equal(t, domain.InOutBounds{}, bound)

	})

	t.Run("Wrong by Admin 1.2", func(t *testing.T) {
		repo.On("UpdateQtyAdminData", mock.Anything).Return(mockData).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.UpdateEntry(input, 5, 2, "admin")

		assert.Equal(t, 200, status)
		assert.Equal(t, mockData, bound)

	})
	t.Run("Wrong by User 1.1", func(t *testing.T) {
		repo.On("UpdateQtyUserData", mock.Anything).Return(emptymockData).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.UpdateEntry(input, 1, 1, "user")

		assert.Equal(t, 404, status)
		assert.Equal(t, domain.InOutBounds{}, bound)

	})

	t.Run("Wrong by User 1.2", func(t *testing.T) {
		repo.On("UpdateQtyUserData", mock.Anything).Return(mockData2).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.UpdateEntry(input, 5, 2, "user")

		assert.Equal(t, 400, status)
		assert.Equal(t, domain.InOutBounds{}, bound)

	})
	t.Run("Wrong by User 1.3", func(t *testing.T) {
		repo.On("UpdateQtyUserData", mock.Anything).Return(mockData).Once()
		useCase := New(repo, validator.New())
		bound, status := useCase.UpdateEntry(input, 5, 2, "user")

		assert.Equal(t, 200, status)
		assert.Equal(t, mockData, bound)

	})
}

func TestDeleteEntry(t *testing.T) {
	repo := new(mocks.InOutBoundData)
	t.Run("Wrong by Admin 1.1", func(t *testing.T) {
		repo.On("DeleteEntryAdminData", mock.Anything).Return("cannot delete data").Once()
		useCase := New(repo, validator.New())
		status := useCase.DeleteEntry(1, 1, "admin")

		assert.Equal(t, 500, status)
	})
	t.Run("Wrong by Admin 1.2", func(t *testing.T) {
		repo.On("DeleteEntryAdminData", mock.Anything).Return("no data deleted").Once()
		useCase := New(repo, validator.New())
		status := useCase.DeleteEntry(1, 1, "admin")

		assert.Equal(t, 404, status)
	})
	t.Run("Wrong by Admin 1.3", func(t *testing.T) {
		repo.On("DeleteEntryAdminData", mock.Anything).Return("").Once()
		useCase := New(repo, validator.New())
		status := useCase.DeleteEntry(1, 1, "admin")

		assert.Equal(t, 204, status)
	})
	t.Run("Wrong by User 1.1", func(t *testing.T) {
		repo.On("DeleteEntryUserData", mock.Anything, mock.Anything).Return("cannot delete data").Once()
		useCase := New(repo, validator.New())
		status := useCase.DeleteEntry(1, 1, "user")

		assert.Equal(t, 500, status)
	})
	t.Run("Wrong by Admin 1.2", func(t *testing.T) {
		repo.On("DeleteEntryUserData", mock.Anything, mock.Anything).Return("no data deleted").Once()
		useCase := New(repo, validator.New())
		status := useCase.DeleteEntry(1, 1, "user")

		assert.Equal(t, 404, status)
	})
	t.Run("Wrong by Admin 1.3", func(t *testing.T) {
		repo.On("DeleteEntryUserData", mock.Anything, mock.Anything).Return("").Once()
		useCase := New(repo, validator.New())
		status := useCase.DeleteEntry(1, 1, "user")

		assert.Equal(t, 204, status)
	})
}
