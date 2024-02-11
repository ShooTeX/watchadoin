package mail

import (
	"fmt"
	"net/smtp"
)

type Mail struct {
	smtp.Auth
	username string
	password string
	smtpHost string
	smtpPort string
}

type MailOptions struct {
	Username string
	Password string
	SmtpHost string
	SmtpPort string
}

func New(o *MailOptions) *Mail {
	return &Mail{
		username: o.Username,
		password: o.Password,
		smtpHost: o.SmtpHost,
		smtpPort: o.SmtpPort,
		Auth:     smtp.PlainAuth("", o.Username, o.Password, o.SmtpHost),
	}
}

func (m *Mail) SendMail(subject, message, from, to string) error {
	message = fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n%s", from, to, subject, message)

	return smtp.SendMail(m.smtpHost+":"+m.smtpPort, m.Auth, from, []string{to}, []byte(message))
}
