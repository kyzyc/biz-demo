package main

import (
	"context"
	"fmt"
	"github.com/kyzyc/biz-demo/app/user/biz/service"
	user "github.com/kyzyc/biz-demo/rpc_gen/kitex_gen/user"
	"strconv"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)
	fmt.Println("---------------\n" + strconv.Itoa(int(resp.UserId)) + "\n---------------------------\n")

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}
