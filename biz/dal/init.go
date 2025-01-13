package dal

import (
	"github.com/kyzyc/biz-demo/app/frontend/biz/dal/mysql"
	"github.com/kyzyc/biz-demo/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
