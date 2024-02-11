package main

import (
	"log"
	"os"

	"github.com/shootex/watchadoin/mail"
)

var (
	from = os.Getenv("MAIL_FROM")
	to   = os.Getenv("MAIL_TO")
)

func main() {
	m := mail.New(&mail.MailOptions{
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
		SmtpHost: os.Getenv("MAIL_SMTP_HOST"),
		SmtpPort: os.Getenv("MAIL_SMTP_PORT"),
	})

	log.Println("Sending mail")
	err := m.SendMail("Oh it's a new mail!", "Hello, World!", from, to)
	if err != nil {
		log.Println("error sending mail:", err)
	} else {
		log.Println("Mail sent!")
	}
}
