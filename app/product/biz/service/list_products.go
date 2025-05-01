package service

import (
	"context"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/product/biz/dal/mysql"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/product/biz/model"
	product "github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	c, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductsResp{}

	if req.CategoryName == "" {
		productQuery := model.NewProductQuery(s.ctx, mysql.DB)
		products, err := productQuery.GetAllProduct()
		if err != nil {
			return nil, err
		}
		for _, p := range products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          int32(p.ID),
				Name:        p.Name,
				Price:       float32(p.Price),
				Description: p.Description,
				Picture:     p.Picture,
			})
		}
		return resp, nil
	}
	//在循环里访问 vl.Products 是成立的，前提是你 用了 Preload 加载了它
	for _, vl := range c {
		for _, p := range vl.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          int32(p.ID),
				Name:        p.Name,
				Price:       float32(p.Price),
				Description: p.Description,
				Picture:     p.Picture,
			})
		}
	}
	return resp, nil
}
