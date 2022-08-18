package helper

import (
	_data "middleman-capstone/feature/orders/data"
	"os"
	"strconv"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func Payment(data _data.OrderPayment) (orderIDGen string, snapResp *snap.Response) {
	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// generate order id
	orderIDGen = strconv.FormatInt(time.Now().Unix(), 10)
	PhoneNumber := strconv.Itoa(data.Phone)
	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderIDGen,
			GrossAmt: int64(data.GrandTotal),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: data.Name,
			Email: data.Email,
			Phone: PhoneNumber,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ = s.CreateTransaction(req)

	return orderIDGen, snapResp
}
