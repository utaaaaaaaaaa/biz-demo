package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/checkout/conf"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/common/clientsuite"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"

	"github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegisterAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}
	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}

}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}

}

func initPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}

}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}
	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}

}
