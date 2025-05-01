package notify

import (
	"github.com/kr/pretty"
	"github.com/utaaaaaaaaaa/biz-demo/gomall/rpc_gen/kitex_gen/email"
)

type NoopEmail struct {
}

func (n *NoopEmail) Send(req *email.EmailReq) error {
	_, err := pretty.Printf("%v\n", req)
	if err != nil {
		return err
	}
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
