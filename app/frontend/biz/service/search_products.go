package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/kyzyc/biz-demo/app/frontend/hertz_gen/frontend/common"
	product "github.com/kyzyc/biz-demo/app/frontend/hertz_gen/frontend/product"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
