package mailer

import (
	"fmt"

	"github.com/matcornic/hermes/v2"
)

// AccountValidEmail 验证 email
func AccountValidEmail(to string, nickname string, token string) error {
	body, err := accountValidEmailBody(to, nickname, token)
	if err != nil {
		return err
	}
	return sendMail([]string{to}, "来自·见周边·的验证信息", body)
}

// accountValidEmailBody 生成邮件体
func accountValidEmailBody(mail, nickname, token string) (string, error) {
	email := hermes.Email{
		Body: hermes.Body{
			Name: nickname,
			Intros: []string{
				"您的邮箱正在用来注册 见周边 的账户。",
			},
			Actions: []hermes.Action{
				{
					Instructions: "点击下面的按钮来验证本邮箱：",
					Button: hermes.Button{
						Color: "#DC4D2F",
						Text:  "验证邮箱",
						Link:  fmt.Sprintf("http://%s/account/valid_email?mail=%s&token=%s", siteHost, mail, token),
					},
				},
			},
		}}

	return her.GenerateHTML(email)
}
