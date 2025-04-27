package dal

import (
	"github.com/utaaaaaaaaaa/biz-demo/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
