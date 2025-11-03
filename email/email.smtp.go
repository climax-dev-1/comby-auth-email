package email

import (
	"net/smtp"
)

type emailSmtp struct {
	options      EmailOptions
	senderEmail  string
	smtpPassword string
	smtpHost     string
	smtpPort     string
}

// Make sure it fullfills interfaces
var _ Email = (*emailSmtp)(nil)

func NewEmailSmtp(senderEmail, smtpPassword, smtpHost, smtpPort string) *emailSmtp {
	e := &emailSmtp{
		options:      EmailOptions{},
		senderEmail:  senderEmail,
		smtpPassword: smtpPassword,
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
	}
	return e
}

func (e *emailSmtp) Init(opts ...EmailOption) error {
	for _, opt := range opts {
		if _, err := opt(&e.options); err != nil {
			return err
		}
	}
	return nil
}

func (e *emailSmtp) SendMail(to []string, subject, message string) error {
	messageBytes := []byte(message)
	auth := smtp.PlainAuth("", e.senderEmail, e.smtpPassword, e.smtpHost)
	if err := smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.senderEmail, to, messageBytes); err != nil {
		return err
	}
	logger.With("email", "smtp").Debug("Email sent", "to", to, "subject", subject)
	return nil
}

func (e *emailSmtp) Options() EmailOptions {
	return e.options
}

func (e *emailSmtp) String() string {
	return "EmailSmtp"
}
