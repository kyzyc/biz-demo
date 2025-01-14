package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	frontendUtils "github.com/kyzyc/biz-demo/frontend/utils"
	"github.com/kyzyc/biz-demo/rpc_gen/kitex_gen/user/userservice"
	"sync"
)

var (
	UserClient userservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		iniUserClient()
	})
}

func iniUserClient() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	frontendUtils.MustHandleError(err)
	if err != nil {
		hlog.Fatal(err)
	}
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}
