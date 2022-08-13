package domain

import (
	"github.com/labstack/echo/v4"
)

type InOutBounds struct {
	ID        int
	IdUser    int
	IdProduct int
	Name      string
	Unit      string
	Qty       int
	Role      string
	Note      string
}

type InOutBoundHandler interface {
	Add() echo.HandlerFunc
	ReadAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type InOutBoundUseCase interface {
	AddEntry(newProduct InOutBounds, id int, role string) (InOutBounds, int)
	ReadEntry(id int, role string) ([]InOutBounds, int)
	UpdateEntry(updatedData InOutBounds, productid, id int, role string) (InOutBounds, int)
	DeleteEntry(productid, id int, role string) int
}

type InOutBoundData interface {
	AddEntryData(newProduct InOutBounds) InOutBounds
	CekUserEntry(newProduct InOutBounds) (cek bool, idcart, qty int)
	CekOwnerEntry(newProduct InOutBounds) (cek bool)
	CekAdminEntry(newProduct InOutBounds) (cek bool, idcart, qty int)
	UpdateQty(idcart int, qty int) InOutBounds
	ReadEntryUserData(id int) []InOutBounds
	ReadEntryAdminData(role string) []InOutBounds
	UpdateEntryAdminData(idproduct int) InOutBounds
	UpdateEntryUserData(idproduct int, id int) InOutBounds
	UpdateQtyUserData(updatedData InOutBounds) InOutBounds
	UpdateQtyAdminData(updatedData InOutBounds) InOutBounds
	DeleteEntryUserData(productid, id int) (row int, err error)
	DeleteEntryAdminData(Productid int) (row int, err error)
}
