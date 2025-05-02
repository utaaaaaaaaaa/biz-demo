package dal

import (
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/email/biz/dal/mysql"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
