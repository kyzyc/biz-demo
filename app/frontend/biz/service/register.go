package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/kyzyc/biz-demo/app/frontend/hertz_gen/frontend/auth"
	common "github.com/kyzyc/biz-demo/app/frontend/hertz_gen/frontend/common"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// TODO user svc api
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	err = session.Save()
	return
}
