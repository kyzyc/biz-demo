package dal

import (
	"github.com/kyzyc/biz-demo/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
