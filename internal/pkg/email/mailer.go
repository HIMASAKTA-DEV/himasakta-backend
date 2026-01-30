package mailer

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type (
	EmailConfig struct {
		Host         string
		Port         int
		SenderName   string
		AuthEmail    string
		AuthPassword string
	}

	Mailer struct {
		emailConfig *EmailConfig
		Body        string
		Error       error
	}
)

func New() Mailer {
	portStr := os.Getenv("SMTP_PORT")
	if portStr == "" {
		// Default SMTP port if not set
		portStr = "587"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic("invalid smtp port")
	}

	emailConfig := &EmailConfig{
		Host:         os.Getenv("SMTP_HOST"),
		Port:         port,
		SenderName:   os.Getenv("SMTP_SENDER_NAME"),
		AuthEmail:    os.Getenv("SMTP_AUTH_EMAIL"),
		AuthPassword: os.Getenv("SMTP_AUTH_PASSWORD"),
	}

	return Mailer{emailConfig, "", nil}
}

func (m Mailer) Send(toEmail, subject string) Mailer {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", m.emailConfig.SenderName)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", m.Body)

	dialer := gomail.NewDialer(
		m.emailConfig.Host,
		m.emailConfig.Port,
		m.emailConfig.AuthEmail,
		m.emailConfig.AuthPassword,
	)

	if err := dialer.DialAndSend(mailer); err != nil {
		m.Error = err
		return m
	}

	return m
}
