package pkg

import (
	"github.com/smartwalle/alipay/v3"
	"hello/ali_pay/bootstrap"
	"time"
)

const (
	webReturnURL  = "http://localhost:8088/return"
	WebTestAmount = "0.01"
)

func WebPageAlipay() (payUrl string, err error) {
	pay := alipay.TradePagePay{}
	pay.ReturnURL = webReturnURL
	pay.Subject = "支付宝网页支付测试"
	pay.OutTradeNo = time.Now().String()
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	pay.TotalAmount = WebTestAmount

	payClient := bootstrap.PayClientInstance()
	url, err := payClient.TradePagePay(pay)
	if err != nil {
		return payUrl, err
	} else {
		payUrl = url.String()
		return payUrl, nil
	}
}
