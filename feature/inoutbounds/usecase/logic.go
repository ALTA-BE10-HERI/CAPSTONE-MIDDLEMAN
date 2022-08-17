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

func (iobuc *inoutboundUseCase) AddEntry(newProduct domain.InOutBounds, id int, role string) (domain.InOutBounds, int) {
	var cart = data.FromIOB(newProduct)
	validError := iobuc.validate.Struct(cart)

	if validError != nil {
		log.Println("Validation error : ", validError)
		return domain.InOutBounds{}, 400
	}

	cart.IdUser = id
	cart.Role = role

	if role != "user" {
		cekowner := iobuc.inoutboundData.CekOwnerAdminEntry(cart.ToIOB())
		if cekowner {
			cek, cartid, cartqty := iobuc.inoutboundData.CekAdminEntry(cart.ToIOB())
			if cek {
				upqty := iobuc.inoutboundData.UpdateQty(cartid, cartqty+1)
				if upqty.ID == 0 {
					log.Println("failed to update data")
					return domain.InOutBounds{}, 500
				}
				return upqty, 201
			} else {
				create := iobuc.inoutboundData.AddEntryData(cart.ToIOB())
				if create.ID == 0 {
					log.Println("error after creating data")
					return domain.InOutBounds{}, 500
				}
				create2 := iobuc.inoutboundData.UpdateEntryAdminData(cart.IdProduct)
				if create2.ID == 0 {
					log.Println("failed to update data")
					return domain.InOutBounds{}, 500
				}
				return create2, 201
			}
		} else {
			return domain.InOutBounds{}, 404
		}
	} else {
		cekowner := iobuc.inoutboundData.CekOwnerEntry(cart.ToIOB())
		if cekowner {
			cek, cartid, cartqty := iobuc.inoutboundData.CekUserEntry(cart.ToIOB())
			if cek {
				upqty := iobuc.inoutboundData.UpdateQty(cartid, cartqty+1)
				if upqty.ID == 0 {
					log.Println("failed to update data")
					return domain.InOutBounds{}, 500
				}
				return upqty, 201
			} else {
				create := iobuc.inoutboundData.AddEntryData(cart.ToIOB())
				if create.ID == 0 {
					log.Println("error after creating data")
					return domain.InOutBounds{}, 500
				}
				create2 := iobuc.inoutboundData.UpdateEntryUserData(cart.IdProduct, cart.IdUser)
				if create2.ID == 0 {
					log.Println("failed to update data")
					return domain.InOutBounds{}, 500
				}
				return create2, 201
			}
		} else {
			return domain.InOutBounds{}, 404
		}
	}
}

func (iobuc *inoutboundUseCase) ReadEntry(id int, role string) ([]domain.InOutBounds, int) {

	if role == "admin" {
		cart := iobuc.inoutboundData.ReadEntryAdminData(role)
		if len(cart) == 0 {
			log.Println("data not found")
			return []domain.InOutBounds{}, 200
		} else {
			return cart, 200
		}
	} else {
		cart := iobuc.inoutboundData.ReadEntryUserData(id)
		if len(cart) == 0 {
			log.Println("data not found")
			return []domain.InOutBounds{}, 200
		} else {
			return cart, 200
		}
	}
}

func (iobuc *inoutboundUseCase) UpdateEntry(updatedData domain.InOutBounds, productid, id int, role string) (domain.InOutBounds, int) {

	if updatedData.Qty == 0 {
		return domain.InOutBounds{}, 400
	}

	updatedData.IdUser = id
	updatedData.IdProduct = productid
	updatedData.Role = role

	if updatedData.Role == "admin" {
		cart := iobuc.inoutboundData.UpdateQtyAdminData(updatedData)
		if cart.ID == 0 {
			log.Println("Empty Data")
			return domain.InOutBounds{}, 404
		}
		return cart, 200
	} else {
		cart := iobuc.inoutboundData.UpdateQtyUserData(updatedData)
		if cart.Note == "insufficient stok" {
			log.Println("insufficient stock")
			return domain.InOutBounds{}, 400
		}
		if cart.ID == 0 {
			log.Println("Empty Data")
			return domain.InOutBounds{}, 404
		}
		return cart, 200
	}
}

func (iobuc *inoutboundUseCase) DeleteEntry(productid, id int, role string) int {

	if role == "admin" {
		err := iobuc.inoutboundData.DeleteEntryAdminData(productid)
		if err == "cannot delete data" {
			log.Println("internal server error")
			return 500
		}
		if err == "no data deleted" {
			log.Println("data not found")
			return 404
		}
		return 204
	} else {
		err := iobuc.inoutboundData.DeleteEntryUserData(productid, id)
		if err == "cannot delete data" {
			log.Println("internal server error")
			return 500
		}
		if err == "no data deleted" {
			log.Println("data not found")
			return 404
		}
		return 204
	}
}
