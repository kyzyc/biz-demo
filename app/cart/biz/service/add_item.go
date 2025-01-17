package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/kyzyc/biz-demo/app/cart/biz/dal/mysql"
	"github.com/kyzyc/biz-demo/app/cart/biz/model"
	cart "github.com/kyzyc/biz-demo/app/cart/kitex_gen/cart"
	"github.com/kyzyc/biz-demo/app/cart/rpc"
	"github.com/kyzyc/biz-demo/rpc_gen/kitex_gen/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	productResp, err := rpc.ProductClient.GetProduct(
		s.ctx, &product.GetProductReq{
			Id: req.Item.ProductId,
		},
	)

	if productResp == nil || productResp.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not found")
	}

	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}

	err = model.AddItem(s.ctx, mysql.DB, cartItem)

	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	return &cart.AddItemResp{}, nil
}
