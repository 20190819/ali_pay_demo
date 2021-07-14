package pkg

import (
	"github.com/smartwalle/alipay/v3"
	"hello/ali_pay/bootstrap"
	"time"
)

const (
	wapReturnURL  = "http://localhost:8088/return"
	wapTestAmount = "0.01"
)

func WapPay() (payUrl string, err error) {
	pay := alipay.TradeWapPay{}
	pay.ReturnURL = wapReturnURL
	pay.Subject = "支付宝网页支付测试"
	pay.OutTradeNo = time.Now().String()
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	pay.TotalAmount = wapTestAmount

	url, err := bootstrap.PayClientInstance().TradeWapPay(pay)

	if err != nil {
		return payUrl, err
	} else {
		payUrl = url.String()
		return payUrl, nil
	}
}
