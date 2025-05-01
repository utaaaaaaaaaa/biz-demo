package dal

import (
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/checkout/biz/dal/mysql"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
