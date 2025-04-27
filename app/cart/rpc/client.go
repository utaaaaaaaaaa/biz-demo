package rpc

import (
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	frontendUtils "github.com/utaaaaaaaaaa/biz-demo/gomall/app/cart/utils"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/frontend/conf"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}
