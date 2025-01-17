package dal

import (
	"github.com/kyzyc/biz-demo/app/cart/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
