package sms

import (
	"testing"

	"github.com/agenewtech/mob-sdk/comm"
)

var client = comm.NewClient("")

func Test_CustomMsg(t *testing.T) {
	req := &CustomMsgRequest{}
	resp := &CustomMsgResponse{}
	err := client.Send(req).Unmarshal(resp)
	t.Log(err)
	t.Logf("%+v", resp)
}

func Test_SMSVerify(t *testing.T) {
	req := &SMSVerifyRequest{
		Phone: "18100000000",
		Zone:  "86",
		Code:  "1256",
	}
	resp := &SMSVerifyResponse{}
	err := client.Send(req).Unmarshal(resp)
	t.Log(err)
	t.Logf("%+v", resp)
}
