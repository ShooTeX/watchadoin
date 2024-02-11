package main

import (
	"os"

	"github.com/shootex/watchadoin/mail"
)

var to = os.Getenv("MAIL_TO")

func main() {
	m := mail.New(&mail.MailOptions{
		From:     os.Getenv("MAIL_FROM"),
		Password: os.Getenv("MAIL_PASSWORD"),
		SmtpHost: os.Getenv("MAIL_SMTP_HOST"),
		SmtpPort: os.Getenv("MAIL_SMTP_PORT"),
	})

	m.SendMail("Hello, World!", to)
}
