package usecase

import (
	"errors"
	"fmt"
	"log"
	"middleman-capstone/domain"
	user "middleman-capstone/feature/users"
	"middleman-capstone/feature/users/data"

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

func (uc *userUseCase) CreateProduct(newProduct domain.ProductUser, id int) int {
	var product = data.FromPU(newProduct)
	validError := uc.validate.Struct(product)

	if validError != nil {
		log.Println("Validation error : ", validError)
		return 400
	}

	product.IdUser = id
	create := uc.userData.CreateProductData(product.ToPU())

	if create.ID == 0 {
		log.Println("error after creating data")
		return 500
	}
	return 201
}

func (uc *userUseCase) ReadAllProduct(id int) ([]domain.ProductUser, int) {
	product := uc.userData.ReadAllProductData(id)
	if len(product) == 0 {
		log.Println("data not found")
		return nil, 404
	}

	return product, 200
}

func (uc *userUseCase) UpdateProduct(updatedData domain.ProductUser, productid, id int) (row int, err error) {
	qry := map[string]interface{}{}

	if updatedData.Name != "" {
		qry["name"] = updatedData.Name
	}

	if updatedData.Unit != "" {
		qry["unit"] = updatedData.Unit
	}

	if updatedData.Stock != 0 {
		qry["stock"] = updatedData.Stock
	}

	if updatedData.Price != 0 {
		qry["price"] = updatedData.Price
	}

	if updatedData.Image != "" {
		qry["image"] = updatedData.Image
	}

	row, err = uc.userData.UpdateProductData(qry, productid, id)

	return row, err
}

func (uc *userUseCase) DeleteProduct(productid, id int) int {
	row, err := uc.userData.DeleteProductData(productid, id)

	if err != nil {
		log.Println("data not found")
		return 404
	}

	if row < 1 {
		log.Println("internal server error")
		return 500
	}

	return 200
}

func (uc *userUseCase) CreateInventory(newRecap domain.InventoryProduct, id int) int {
	var product = data.FromIP(newRecap)
	validError := uc.validate.Struct(product)

	if validError != nil {
		log.Println("Validation error : ", validError)
		return 400
	}

	product.IdUser = id
	create := uc.userData.CreateInventoryData(product.ToIP())

	if create.ID == 0 {
		log.Println("error after creating data")
		return 500
	}
	return 201
}
