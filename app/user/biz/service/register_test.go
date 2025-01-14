package service

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/kyzyc/biz-demo/app/user/biz/dal/mysql"
	user "github.com/kyzyc/biz-demo/rpc_gen/kitex_gen/user"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "demo@admin.com",
		Password:        "FJOSJDOJASD",
		PasswordConfirm: "FJOSJDOJASD",
	}
	resp, err := s.Run(req)

	t.Logf("err: %v", err)
	t.Logf("resp.UserId: %v", resp.UserId)

	// todo: edit your unit test

}
