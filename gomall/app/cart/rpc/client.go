package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/cart/conf"
	cartutils "github.com/utaaaaaaaaaa/biz-demo/gomall/app/cart/utils"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/common/clientsuite"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegisterAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
