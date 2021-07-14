package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	AppId        string
	AliPublicKey string
	PrivateKey   string
	IsProd       bool
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(".env 配置文件未找到")
			return
		}
	} else {
		AppId = viper.GetString("APPID")
		AliPublicKey = viper.GetString("PUBLIC_KEY")
		PrivateKey = viper.GetString("PRIVATE_KEY")
		IsProd = viper.GetBool("PROD")
		fmt.Println("===加载配置文件成功===")
	}
}
