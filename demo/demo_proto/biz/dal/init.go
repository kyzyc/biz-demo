package dal

import (
	"github.com/kyzyc/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/kyzyc/biz-demo/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
