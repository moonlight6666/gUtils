package gUtils

import (
	"gopkg.in/gomail.v2"
)

type MailOptions struct {
	MailHost           string
	MailPort           int
	MailUser           string   // 发件人
	MailUserFormatName string   // 发件人
	MailPass           string   // 发件人密码
	MailTo             []string // 收件人
	Subject            string   // 邮件主题
	Body               string   // 邮件内容
}

func SendMail(o *MailOptions) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(o.MailUser, o.MailUserFormatName))
	m.SetHeader("To", o.MailTo...)    //发送给多个用户
	m.SetHeader("Subject", o.Subject) //设置邮件主题
	m.SetBody("text/html", o.Body)    //设置邮件正文
	d := gomail.NewDialer(o.MailHost, o.MailPort, o.MailUser, o.MailPass)

	return d.DialAndSend(m)
}
