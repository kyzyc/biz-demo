package dal

import (
	"github.com/kyzyc/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/kyzyc/biz-demo/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
