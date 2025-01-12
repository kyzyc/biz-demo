package home

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/kyzyc/biz-demo/gomall/app/frontend/biz/service"
	"github.com/kyzyc/biz-demo/gomall/app/frontend/biz/utils"
	"github.com/kyzyc/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := map[string]any{}
	resp, err = service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "home", utils.WrapResponse(ctx, c, resp))

	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
