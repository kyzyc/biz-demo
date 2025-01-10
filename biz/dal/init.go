package dal

import (
	"github.com/kyzyc/biz-demo/gomall/app/frontend/biz/dal/mysql"
	"github.com/kyzyc/biz-demo/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
