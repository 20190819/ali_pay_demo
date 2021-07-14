package bootstrap

import (
	"github.com/smartwalle/alipay/v3"
	"hello/ali_pay/config"
	"sync"
)

var once sync.Once
var client *alipay.Client
var publicKey = config.AliPublicKey

func init() {
	client = PayClientInstance()
}

func PayClientInstance() *alipay.Client {
	once.Do(func() {
		client, _ = alipay.New(config.AppId, config.PrivateKey, config.IsProd)
		_=client.LoadAliPayPublicKey(publicKey)
	})

	return client
}
