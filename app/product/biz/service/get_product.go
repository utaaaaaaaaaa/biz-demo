package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/product/biz/dal/mysql"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/product/biz/model"
	product "github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := productQuery.GetProductById(int(req.Id))
	if err != nil {
		return nil, err

	}
	return &product.GetProductResp{
		Product: &product.Product{
			Id:          int32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		},
	}, nil
}
