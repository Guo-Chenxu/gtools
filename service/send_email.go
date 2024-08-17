package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"gtools/biz/model/gtools"
	"gtools/consts"
	"net/smtp"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jordan-wright/email"
)

func SendEmail(ctx context.Context, req *gtools.SendEmailReq) *consts.BizCode {
	auth := smtp.PlainAuth("", req.From, req.Secret, req.Host)
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s<%s>", req.Nickname, req.From)
	e.To = req.To
	e.Subject = req.Subject
	e.HTML = []byte(req.Body)
	hostAddr := fmt.Sprintf("%s:%d", req.Host, req.Port)

	var err error
	if req.Ssl {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: req.Host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	if err != nil {
		hlog.CtxInfof(ctx, "send email failed, err: %v", err)
		return &consts.BizCode{consts.SendEmailError.Code, consts.SendEmailError.Msg}
	}
	return nil
}
