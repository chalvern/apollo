package mailer

import (
	"gopkg.in/gomail.v2"
)

func sendMail(mailTo []string, subject string, body string) error {

	m := gomail.NewMessage()
	// 添加别名
	m.SetHeader("From", m.FormatAddress(emailAddress, emailAliasName))
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(
		emailHost, emailHostPort, emailAddress, emailPassword)
	err := d.DialAndSend(m)
	return err
}
