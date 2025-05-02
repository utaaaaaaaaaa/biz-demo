package email

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/email/infra/mq"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/app/email/infra/notify"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/email"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req) //数据为protobuf形式，通过protobuf反序列化
		if err != nil {
			klog.Error(err)
			return
		}

		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
	})
	if err != nil {
		panic(err)
	}

	//服务退出时取消订阅
	server.RegisterShutdownHook(func() {
		sub.Unsubscribe()
		mq.Nc.Close()
	})
}
