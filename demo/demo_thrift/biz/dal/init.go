package dal

import (
	"github.com/kyzyc/biz-demo/demo/demo_thrift/biz/dal/mysql"
	"github.com/kyzyc/biz-demo/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
