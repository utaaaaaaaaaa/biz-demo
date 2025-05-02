package consumer

import "github.com/utaaaaaaaaaa/biz-demo/gomall/app/email/biz/consumer/email"

// 可能有多个consumer，一起启动
func Init() {
	email.ConsumerInit()
}
