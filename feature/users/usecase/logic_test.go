package usecase

import (
	"errors"
	"middleman-capstone/domain"
	"middleman-capstone/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUser(t *testing.T) {
	repo := new(mocks.MockUser)
	insertData := domain.User{
		ID:       1,
		Name:     "ivan",
		Email:    "ivan@gmail.com",
		Password: "123",
		Phone:    "081217076500",
		Address:  "malang",
		Role:     "user",
	}
	falseData := domain.User{
		ID:       0,
		Name:     "",
		Email:    "",
		Password: "",
		Phone:    "",
		Address:  "",
		Role:     "",
	}
	outputData := 1

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(outputData, nil).Once()

		srv := New(repo, validator.New())

		res, err := srv.AddUser(insertData)
		assert.NoError(t, err)
		assert.Equal(t, outputData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Fill Error", func(t *testing.T) {
		repo.On("Fill Error", mock.Anything).Return(-1, errors.New("please make sure all fields are filled in correctly")).Once()

		srv := New(repo, validator.New())

		res, err := srv.AddUser(falseData)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
	})

}
func TestDeleteCase(t *testing.T) {
	repo := new(mocks.MockUser)
	t.Run("success delete", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything).Return(1, nil)

		srv := New(repo, validator.New())
		delete, err := srv.DeleteCase(1)

		assert.NoError(t, err)
		assert.Equal(t, 1, delete)
	})
}

func TestUpdateCase(t *testing.T) {
	insertData := domain.User{
		Name:     "user",
		Email:    "user@gmail.com",
		Password: "123",
		Phone:    "081217076500",
		Address:  "malang",
	}

	repo := new(mocks.MockUser)

	t.Run("succes update", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, 1).Return(1, nil)
		srv := New(repo, validator.New())
		update, err := srv.UpdateCase(insertData, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, update)

	})
}

func TestGetProfile(t *testing.T) {
	inserData := domain.User{
		Name:    "ivan",
		Email:   "ivan@gmail.com",
		Phone:   "081217076500",
		Address: "malang",
		Role:    "user",
	}

	repo := new(mocks.MockUser)

	t.Run("data found", func(t *testing.T) {
		repo.On("GetSpecific", 1).Return(inserData, nil).Once()

		srv := New(repo, validator.New())
		myprofile, err := srv.GetProfile(1)
		assert.NoError(t, err)
		assert.Equal(t, inserData, myprofile)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("GetSpecific", 0).Return(inserData, nil).Once()

		srv := New(repo, validator.New())
		myprofile, err := srv.GetProfile(0)
		assert.NoError(t, err)
		assert.Equal(t, inserData, myprofile)

	})
}
