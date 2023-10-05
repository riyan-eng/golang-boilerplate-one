package util

import (
	"fmt"
	"net/smtp"

	"github.com/riyan-eng/golang-boilerplate-one/env"
	"github.com/jordan-wright/email"
)

type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name, fromEmailAddress, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (s *GmailSender) SendEmail(subject string, content string, to []string, cc []string, bcc []string, attacthFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", s.name, s.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	smtpAuth := smtp.PlainAuth("", s.fromEmailAddress, s.fromEmailPassword, env.SMTP_HOST)
	return e.Send(env.SMTP_HOST+":"+env.SMTP_PORT, smtpAuth)
}
