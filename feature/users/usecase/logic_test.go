package usecase

import (
	"errors"
	"middleman-capstone/domain"
	"middleman-capstone/domain/mocks"
	user "middleman-capstone/feature/users"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddUser(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := domain.User{ID: 1, Name: "ivan", Email: "ivan@gmail.com", Password: "123", Phone: "081217076500", Address: "malang", Role: "user"}
	falseData := domain.User{ID: 0, Name: "", Email: "", Password: "", Phone: "", Address: "", Role: ""}
	falseData3 := domain.User{ID: 0, Name: "", Email: "ivan@gmail.com", Password: "", Phone: "", Address: "", Role: ""}
	outputData := 1

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(outputData, nil).Once()

		srv := New(repo, validator.New())

		res, err := srv.AddUser(insertData)
		assert.NoError(t, err)
		assert.Equal(t, outputData, res)
		repo.AssertExpectations(t)
	})

	t.Run("your format email is false", func(t *testing.T) {
		repo.On("MatchString", mock.Anything).Return(false).Once()
		srv := New(repo, validator.New())
		res, err := srv.AddUser(falseData)
		assert.Equal(t, 400, res)
		assert.NotNil(t, err)
	})

	t.Run("Fill Error", func(t *testing.T) {
		repo.On("MatchString", mock.Anything).Return(true).Once()
		srv := New(repo, validator.New())
		res, err := srv.AddUser(falseData3)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
	})

}

func TestLogin(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := user.LoginModel{Email: "ivan@gmail.com", Password: "123"}
	result := map[string]interface{}{}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("LoginData", mock.Anything).Return(result, nil).Once()

		srv := New(repo, validator.New())

		res, err := srv.Login(insertData)
		assert.NoError(t, err)
		assert.Equal(t, result, res)
		repo.AssertExpectations(t)
	})
}

func TestGetProfile(t *testing.T) {
	inserData := domain.User{Name: "ivan", Email: "ivan@gmail.com", Phone: "081217076500", Address: "malang", Role: "user"}
	repo := new(mocks.UserData)

	t.Run("data found", func(t *testing.T) {
		repo.On("GetSpecific", mock.Anything).Return(inserData, gorm.ErrRecordNotFound).Once()
		srv := New(repo, validator.New())
		myprofile, err := srv.GetProfile(1)
		assert.Equal(t, errors.New("data not found"), err)
		assert.Equal(t, domain.User{}, myprofile)
	})

	t.Run("data found", func(t *testing.T) {
		repo.On("GetSpecific", mock.Anything).Return(inserData, errors.New("error")).Once()
		srv := New(repo, validator.New())
		myprofile, err := srv.GetProfile(1)
		assert.Equal(t, errors.New("server error"), err)
		assert.Equal(t, domain.User{}, myprofile)
	})

	t.Run("succes", func(t *testing.T) {
		repo.On("GetSpecific", mock.Anything).Return(inserData, nil).Once()

		srv := New(repo, validator.New())
		myprofile, err := srv.GetProfile(1)
		assert.NoError(t, err)
		assert.Equal(t, inserData, myprofile)

	})
}

func TestDeleteCase(t *testing.T) {
	repo := new(mocks.UserData)
	t.Run("success delete", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything).Return(1, nil)

		srv := New(repo, validator.New())
		delete, err := srv.DeleteCase(1)

		assert.NoError(t, err)
		assert.Equal(t, 1, delete)
	})
}

func TestUpdateCase(t *testing.T) {
	insertData := domain.User{Name: "user", Email: "user@gmail.com", Password: "123", Phone: "081217076500", Address: "malang"}
	insertData2 := domain.User{Name: "", Email: "", Password: "", Phone: "", Address: ""}
	insertData3 := domain.User{Name: "Vanili", Email: "", Password: "123", Phone: "", Address: "jalanjalan"}
	repo := new(mocks.UserData)

	t.Run("succes update", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, 1).Return(1, nil)
		srv := New(repo, validator.New())
		update, err := srv.UpdateCase(insertData, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, update)

	})

	t.Run("succes update", func(t *testing.T) {
		srv := New(repo, validator.New())
		update, err := srv.UpdateCase(insertData2, 1)
		assert.NoError(t, err)
		assert.Equal(t, 404, update)

	})

	t.Run("succes update", func(t *testing.T) {
		repo.On("MatchString", mock.Anything).Return(false).Once()
		srv := New(repo, validator.New())
		update, err := srv.UpdateCase(insertData3, 1)
		assert.Equal(t, errors.New("your format email is false"), err)
		assert.Equal(t, 400, update)

	})
}
