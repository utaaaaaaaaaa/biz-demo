package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	product "github.com/utaaaaaaaaaa/biz-demo/gomall/app/frontend/hertz_gen/frontend/product"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{Query: req.Query})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"items": p.Results,
		"query": req.Query,
	}, nil
}
