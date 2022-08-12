package usecase

import (
	"log"
	"middleman-capstone/domain"
	"middleman-capstone/feature/inoutbounds/data"

	"github.com/go-playground/validator"
)

type inoutboundUseCase struct {
	inoutboundData domain.InOutBoundData
	validate       *validator.Validate
}

func New(iobd domain.InOutBoundData, v *validator.Validate) domain.InOutBoundUseCase {
	return &inoutboundUseCase{
		inoutboundData: iobd,
		validate:       v,
	}
}

func (iobuc *inoutboundUseCase) AddEntry(newProduct domain.InOutBounds, id int) int {
	var cart = data.FromIOB(newProduct)
	validError := iobuc.validate.Struct(cart)

	if validError != nil {
		log.Println("Validation error : ", validError)
		return 400
	}

	cart.IdUser = id

	cek, cartid, cartqty := iobuc.inoutboundData.CekEntry(cart.ToIOB())

	if cek {
		row, err := iobuc.inoutboundData.UpdateQty(cartid, cartqty+1)
		if row == 0 || err != nil {
			log.Println("failed to update data")
			return 500
		}
	} else {
		create := iobuc.inoutboundData.AddEntryData(cart.ToIOB())
		if create.ID == 0 {
			log.Println("error after creating data")
			return 500
		}
	}

	return 201
}
