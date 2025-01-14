package dal

import (
	"github.com/kyzyc/biz-demo/app/product/biz/dal/mysql"
	"github.com/kyzyc/biz-demo/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
