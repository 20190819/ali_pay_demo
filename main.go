package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	_ "hello/ali_pay/bootstrap"
	_ "hello/ali_pay/config"
	"hello/ali_pay/pkg"
)

func main() {
	// 网页支付
	payPageUrl, err := pkg.WebPageAlipay()
	if err != nil {
		logrus.Error(err)
	} else {
		fmt.Println("payWapUrl>>> ", payPageUrl)
	}
	//手机端支付
	payWapUrl, err := pkg.WapPay()
	if err != nil {
		logrus.Error(err)
	} else {
		fmt.Println("payWapUrl>>> ", payWapUrl)
	}

}
