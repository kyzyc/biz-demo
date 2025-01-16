package service

import (
	"context"
	"github.com/kyzyc/biz-demo/app/product/biz/dal/mysql"
	"github.com/kyzyc/biz-demo/app/product/biz/model"
	product "github.com/kyzyc/biz-demo/app/product/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)

	products, err := productQuery.SearchProducts(req.Query)
	var results []*product.Product

	for _, v := range products {
		results = append(results, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Picture:     v.Picture,
		})
	}

	return &product.SearchProductsResp{Results: results}, err
}
