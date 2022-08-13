package usecase

import (
	"errors"
	"fmt"
	"log"
	"middleman-capstone/domain"
	user "middleman-capstone/feature/users"
	"regexp"

	_bcrypt "golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type userUseCase struct {
	userData domain.UserData
	validate *validator.Validate
}

func New(ud domain.UserData, v *validator.Validate) domain.UserUseCase {
	return &userUseCase{
		userData: ud,
		validate: v,
	}
}

func (uc *userUseCase) AddUser(newUser domain.User) (row int, err error) {
	emailFormat := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	cekEmail := emailFormat.MatchString(newUser.Email)
	if cekEmail == false {
		return 400, errors.New("your format email is false")
	}
	phoneFormat := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	cekPhone := phoneFormat.MatchString(newUser.Phone)
	if cekPhone == false {
		log.Println("cek phone", newUser.Phone)
		return 401, errors.New("your format phone is false")
	}

	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" || newUser.Phone == "" || newUser.Address == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}

	row, err = uc.userData.Insert(newUser)
	return row, err
}

func (uc *userUseCase) Login(authData user.LoginModel) (data map[string]interface{}, err error) {
	data, err = uc.userData.LoginData(authData)
	return data, err
}

func (uc *userUseCase) GetProfile(id int) (domain.User, error) {
	data, err := uc.userData.GetSpecific(id)

	if err != nil {
		log.Println("Use case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, errors.New("data not found")
		} else {
			return domain.User{}, errors.New("server error")
		}
	}

	return data, nil
}

func (uc *userUseCase) DeleteCase(userID int) (row int, err error) {
	row, err = uc.userData.DeleteData(userID)
	return row, err
}

func (uc *userUseCase) UpdateCase(input domain.User, idFromToken int) (row int, err error) {
	userReq := map[string]interface{}{}
	if input.Name != "" {
		userReq["name"] = input.Name
	}
	if input.Email != "" {
		userReq["email"] = input.Email
	}
	if input.Phone != "" {
		userReq["phone"] = input.Phone
	}
	if input.Address != "" {
		userReq["address"] = input.Address
	}
	if input.Password != "" {
		passwordHashed, errorHash := _bcrypt.GenerateFromPassword([]byte(input.Password), 10)
		if errorHash != nil {
			fmt.Println("Error hash", errorHash.Error())
		}
		userReq["password"] = string(passwordHashed)
	}
	row, err = uc.userData.UpdateData(userReq, idFromToken)
	return row, err
}
