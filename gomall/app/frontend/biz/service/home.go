package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/utaaaaaaaaaa/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	ctx := h.Context
	p, err := rpc.ProductClient.ListProducts(ctx, &product.ListProductsReq{})
	if err != nil {
		klog.Error(err)
	}
	var cartNum int
	return utils.H{
		"title":    "Hot sale",
		"cart_num": cartNum,
		"items":    p.Products,
	}, nil
}
