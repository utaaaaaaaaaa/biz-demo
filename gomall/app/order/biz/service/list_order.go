package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/order/biz/model"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	list, err := model.ListOrder(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(500001, err.Error())
	}
	var orders []*order.Order
	for _, v := range list {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  oi.Quantity,
				},
				Cost: oi.Cost,
			})
		}

		orders = append(orders, &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			CreatedAt:    int32(v.CreatedAt.Unix()),
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				City:          v.Consignee.City,
				State:         v.Consignee.State,
				Country:       v.Consignee.Country,
				ZipCode:       v.Consignee.ZipCode,
			},
			OrderItems: items,
		})
	}

	resp = &order.ListOrderResp{
		Orders: orders,
	}
	return
}
