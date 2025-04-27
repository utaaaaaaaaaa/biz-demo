package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/demo/demo_proto/biz/model"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	dal.Init()
	//insert
	//mysql.DB.Create(&model.User{Email: "demo@01example.com", Password: "123456"})
	//mysql.DB.Create(&model.User{Email: "demo02@example.com", Password: "123456"})
	//update
	mysql.DB.Model(&model.User{}).Where("email like ?", "%01%").Update("password", "1234566")
	//select
	var row = model.User{}
	mysql.DB.Model(&model.User{}).Where("email like ?", "demo%").First(&row)
	fmt.Printf("\n%v", row)
	//delete
	mysql.DB.Where("email like ?", "%02%").Delete(&model.User{})
}
