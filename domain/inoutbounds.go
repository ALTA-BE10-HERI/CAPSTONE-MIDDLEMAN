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
}

type InOutBoundHandler interface {
	Add() echo.HandlerFunc
}

type InOutBoundUseCase interface {
	AddEntry(newProduct InOutBounds, id int) int
}

type InOutBoundData interface {
	AddEntryData(newProduct InOutBounds) InOutBounds
	CekEntry(newProduct InOutBounds) (cek bool, idcart, qty int)
	UpdateQty(idcart int, qty int) (row int, err error)
}
