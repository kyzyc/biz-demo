package service

import (
	"context"
	"github.com/hertz-contrib/sessions"
	"github.com/kyzyc/biz-demo/app/frontend/infra/rpc"
	"github.com/kyzyc/biz-demo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/kyzyc/biz-demo/app/frontend/hertz_gen/frontend/auth"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// TODO USER SVC API
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	err = session.Save()
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return redirect, nil
}
