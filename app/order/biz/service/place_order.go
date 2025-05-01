package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/order/biz/model"
	order "github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	if len(req.OrderItems) == 0 {
		err = kerrors.NewGRPCBizStatusError(500001, "item is empty")
		return nil, err
	}
	//涉及到两张表操作，使用事务
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()

		o := model.Order{
			OrderId:      orderId.String(),
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}

		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.StreetAddress
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.Country = a.Country
			o.Consignee.ZipCode = a.ZipCode
		}

		err = tx.Create(&o).Error
		if err != nil {
			return err
		}

		var items []model.OrderItem
		for _, v := range req.OrderItems {
			items = append(items, model.OrderItem{
				ProductId:    v.Item.ProductId,
				OrderIdRefer: orderId.String(),
				Quantity:     v.Item.Quantity,
				Cost:         v.Cost,
			})
		}

		err = tx.Create(&items).Error
		if err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}
