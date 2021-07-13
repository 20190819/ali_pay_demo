package pkg

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"hello/ali_pay/bootstrap"
	"time"
)

const (
	webReturnURL  = "http://localhost:8088/return"
	WebTestAmount = "0.01"
)

func WebPageAlipay() (payUrl string) {
	pay := alipay.TradePagePay{}
	pay.ReturnURL = webReturnURL
	pay.Subject = "支付宝网页支付测试"
	pay.OutTradeNo = time.Now().String()
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	pay.TotalAmount = WebTestAmount

	url, err := bootstrap.PayClientInstance().TradePagePay(pay)
	if err != nil {
		fmt.Println(err)
		return
	}

	payUrl = url.String()
	return
}

func handle() {
	payUrl := WebPageAlipay()
	fmt.Printf("ali web page pay url == %s", payUrl)
}
