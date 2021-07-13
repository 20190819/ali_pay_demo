package main

import (
	"fmt"
	"github.com/spf13/viper"
	_ "hello/ali_pay/config"
)

func main() {
	viper.SetConfigName(".env")
	//viper.SetConfigType("")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("配置文件未找到")
			return
		}
	} else {
		fmt.Println("appid=", viper.GetString("APPID"))
	}

	//return
	//// 网页支付
	//payUrl := pkg.WebPageAlipay()
	//// 手机端支付
	////payUrl := pkg.WapPay()
	//fmt.Println("payUrl", payUrl)
}
