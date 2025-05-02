package dal

import (
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/user/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
