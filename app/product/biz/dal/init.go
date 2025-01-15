package dal

import (
	"github.com/kyzyc/biz-demo/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
