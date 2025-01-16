package dal

import (
	"github.com/kyzyc/biz-demo/app/cart/biz/dal/mysql"
	"github.com/kyzyc/biz-demo/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
