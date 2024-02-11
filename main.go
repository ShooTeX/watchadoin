package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/playwright-community/playwright-go"
	"github.com/robfig/cron"
	"github.com/shootex/watchadoin/checker"
	"github.com/shootex/watchadoin/mail"
)

var (
	from     = os.Getenv("MAIL_FROM")
	to       = os.Getenv("MAIL_TO")
	url      = os.Getenv("URL")
	selector = os.Getenv("SELECTOR")
)

func main() {
	if err := playwright.Install(); err != nil {
		log.Fatal("could not install playwright:", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		log.Fatal("could not run playwright:", err)
	}

	m := mail.New(&mail.MailOptions{
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
		SmtpHost: os.Getenv("MAIL_SMTP_HOST"),
		SmtpPort: os.Getenv("MAIL_SMTP_PORT"),
	})

	pc := checker.New(pw, url, selector)
	slog.Info("Running first checker")
	if err := runChecker(pc, m); err != nil {
		log.Fatal("error running first checker:", err)
	}

	c := cron.New()

	c.AddFunc("@hourly", func() {
		if err := runChecker(pc, m); err != nil {
			slog.Error(err.Error())
		}
	})

	c.Start()
	defer c.Stop()

	select {}
}

func runChecker(c *checker.Checker, m *mail.Mail) error {
	slog.Info("Checking")
	checkerResponse, err := c.IsSame()
	if err != nil {
		return err
	}

	if !checkerResponse.IsSame {
		slog.Info("Value has changed!", *checkerResponse.OldValue, *checkerResponse.NewValue)
		slog.Info("Sending mail")
		if err := m.SendMail(
			"Changes detected on "+url, "Old value: "+*checkerResponse.OldValue+"\nNew value: "+*checkerResponse.NewValue,
			from,
			to,
		); err != nil {
			return err
		}

		slog.Info("Mail sent!")
		return nil
	}

	slog.Info("Value is the same")
	return nil
}
