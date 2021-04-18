package email

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}
type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

//NewEmail 实例化
func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

//SendMail 发送邮件
func (e *Email) SendMail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("Form", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetHeader("text/html", body)
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}
