package utils

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// var cfg = config.LoadConfig()

func SendEmail(subject, fromEmail, name, message string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USERNAME")
	smtpPass := os.Getenv("SMTP_PASSWORD")
	receiverEmail := os.Getenv("RECEIVER_EMAIL")

	// Convert SMTP port from string to int
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return fmt.Errorf("invalid SMTP_PORT: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", receiverEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", fmt.Sprintf("From: %s\nName: %s\n\nMessage:\n%s", fromEmail, name, message))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)

	return d.DialAndSend(m)
}
