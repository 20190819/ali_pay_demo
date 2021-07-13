package pkg

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"hello/ali_pay/bootstrap"
	"time"
)

const (
	wapReturnURL  = "http://localhost:8088/return"
	wapTestAmount = "0.01"
)

func WapPay() (payUrl string) {
	pay := alipay.TradeWapPay{}
	pay.ReturnURL = wapReturnURL
	pay.Subject = "支付宝网页支付测试"
	pay.OutTradeNo = time.Now().String()
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	pay.TotalAmount = wapTestAmount

	url, err := bootstrap.PayClientInstance().TradeWapPay(pay)

	if err != nil {
		fmt.Println(err)
		return
	}

	payUrl = url.String()
	return
}
