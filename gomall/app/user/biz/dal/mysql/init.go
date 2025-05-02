package mysql

import (
	"fmt"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/user/biz/model"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/user/conf"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	//	dsn := fmt.Sprintf("%s:%s@tcp(%s:13306)/user?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	//主要用来让数据库表和字段同步
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
