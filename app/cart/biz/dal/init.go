package dal

import (
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/cart/biz/dal/mysql"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
