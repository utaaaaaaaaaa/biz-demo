package dal

import (
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/frontend/biz/dal/mysql"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
