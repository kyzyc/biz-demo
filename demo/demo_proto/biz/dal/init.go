package dal

import (
	"github.com/kyzyc/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
