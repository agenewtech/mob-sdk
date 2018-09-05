package sms

import (
	"bytes"

	"github.com/agenewtech/mob-sdk/comm"
)

const (
	SMSEndPoint  = "https://webapi.sms.mob.com"
	SMSVerifyURL = SMSEndPoint + "/sms/verify"
	CustomMsgURL = SMSEndPoint + "/custom/msg"
)

type CustomMsgRequest struct {
	TemplateCode string
	Zone         string
	Phone        string
	Name         string
	Address      string
	NeedID       int
}

func (r *CustomMsgRequest) MarshalURL(c *comm.Client, buf *bytes.Buffer) error {
	buf.WriteString(CustomMsgURL)
	buf.WriteString("?appKey=")
	buf.WriteString(c.AppKey)
	buf.WriteString("&templateCode=")
	buf.WriteString(r.TemplateCode)
	buf.WriteString("&zone=")
	buf.WriteString(r.Zone)
	buf.WriteString("&phone=")
	buf.WriteString(r.Phone)
	buf.WriteString("&name=")
	buf.WriteString(r.Name)
	buf.WriteString("&address=")
	buf.WriteString(r.Address)
	if r.NeedID == 1 {
		buf.WriteString("&needId=1")
	}
	return nil
}

type CustomMsgResponse struct {
	Status int
	Res    struct {
		SMSID string
	}
}

type SMSVerifyRequest struct {
	Phone string
	Zone  string
	Code  string
}

func (r *SMSVerifyRequest) MarshalURL(c *comm.Client, buf *bytes.Buffer) error {
	buf.WriteString(CustomMsgURL)
	buf.WriteString("?appKey=")
	buf.WriteString(c.AppKey)
	buf.WriteString("&zone=")
	buf.WriteString(r.Zone)
	buf.WriteString("&phone=")
	buf.WriteString(r.Phone)
	buf.WriteString("&code=")
	buf.WriteString(r.Code)
	return nil
}

type SMSVerifyResponse struct {
	Status int
}
