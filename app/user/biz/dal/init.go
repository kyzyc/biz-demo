package dal

import (
	"github.com/kyzyc/biz-demo/app/user/biz/dal/mysql"
	"github.com/kyzyc/biz-demo/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
