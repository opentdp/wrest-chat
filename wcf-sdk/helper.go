package wcf

import (
	"google.golang.org/protobuf/proto"
)

type cmdMSG struct {
	*Request
}

func (c *cmdMSG) build() []byte {
	marshal, _ := proto.Marshal(c)
	return marshal
}

func genFunReq(fun Functions) *cmdMSG {
	return &cmdMSG{
		&Request{Func: fun, Msg: nil},
	}
}

func NewWCF(addr string) (*Client, error) {
	client := &Client{
		server: "tcp://" + addr,
	}
	return client, client.dial()
}
