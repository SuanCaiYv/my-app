package service

import (
	"fmt"
	config2 "github.com/SuanCaiYv/my-app-backend/config"
	"net/smtp"
	"sync"
)

type EmailService interface {
	SendVerCode(target, verCode string) error
}

type EmailServiceStruct struct {
}

var (
	emailServiceOnce sync.Once
	emailService     *EmailServiceStruct
)

func NewEmailService() EmailService {
	newEmailService()
	return emailService
}

func newEmailService() {
	emailServiceOnce.Do(func() {
		emailService = &EmailServiceStruct{}
	})
}

func (e *EmailServiceStruct) SendVerCode(target, verCode string) error {
	contentType := "Content-Type: text/html; charset=UTF-8"
	body := `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="utf-8">
			<title></title>
		</head>
		<body>` + "【CodeWithBuff】您的验证码是：<strong>" + verCode + "</strong>，两分钟内有效。" +
		`</body>
		</html>`
	msg := []byte("To: " + target + "\r\n" +
		"Subject: Verification Code\r\n" +
		contentType + "\r\n\r\n" +
		body)
	config := config2.ApplicationConfiguration()
	auth := smtp.PlainAuth("", config2.ApplicationConfiguration().Email.Sender, config.Email.Credential, config.Email.Host)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", config.Email.Host, config.Email.Port), auth, config.Email.Sender, []string{target}, msg)
	return err
}
