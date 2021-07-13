package bootstrap

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"hello/ali_pay/config"
	"sync"
)

var once sync.Once
var Client *alipay.Client

func init() {
	err := PayClientInstance().LoadAliPayPublicKey(config.AliPublicKey)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func PayClientInstance() *alipay.Client {
	once.Do(func() {
		Client, _ = alipay.New(config.AppId, config.PrivateKey, config.IsProd)
	})
	return Client
}
