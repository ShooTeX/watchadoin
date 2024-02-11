package mail

import (
	"log"
	"net/smtp"
)

type Mail struct {
	smtp.Auth
	from     string
	password string
	smtpHost string
	smtpPort string
}

type MailOptions struct {
	From     string
	Password string
	SmtpHost string
	SmtpPort string
}

func New(o *MailOptions) *Mail {
	return &Mail{
		from:     o.From,
		password: o.Password,
		smtpHost: o.SmtpHost,
		smtpPort: o.SmtpPort,
		Auth:     smtp.PlainAuth("", o.From, o.Password, o.SmtpHost),
	}
}

func (m *Mail) SendMail(message, to string) {
	log.Printf("sending mail to %s", to)
}
