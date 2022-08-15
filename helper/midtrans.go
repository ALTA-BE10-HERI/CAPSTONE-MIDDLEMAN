package helper

import (
	"os"
	"strconv"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type DetailPayment struct {
	Name         string
	Email        string
	Handphone    string
	TotalPayment int
}

func Payment(dataUser DetailPayment) (orderID string, snapResp *snap.Response) {
	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// generate order id
	orderIDGen := strconv.FormatInt(time.Now().Unix(), 10)

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderIDGen,
			GrossAmt: int64(dataUser.TotalPayment),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataUser.Name,
			Email: dataUser.Email,
			Phone: dataUser.Handphone,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ = s.CreateTransaction(req)

	return orderIDGen, snapResp
}
