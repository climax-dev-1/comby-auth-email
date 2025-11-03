package email

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type emailSendgrid struct {
	options        EmailOptions
	sendGridApiKey string
	senderEmail    string
	senderName     string
}

// Make sure it fullfills interfaces
var _ Email = (*emailSendgrid)(nil)

func NewEmailSendgrid(
	sendGridApiKey,
	senderEmail,
	senderName string) *emailSendgrid {
	e := &emailSendgrid{
		options:        EmailOptions{},
		sendGridApiKey: sendGridApiKey,
		senderEmail:    senderEmail,
		senderName:     senderName,
	}
	return e
}

func (e *emailSendgrid) Init(opts ...EmailOption) error {
	for _, opt := range opts {
		if _, err := opt(&e.options); err != nil {
			return err
		}
	}
	return nil
}

func (e *emailSendgrid) SendMail(to []string, subject, message string) error {
	// check
	if len(to) <= 0 {
		return fmt.Errorf("to is invalid: %s", to)
	}
	if len(subject) <= 0 {
		return fmt.Errorf("subject is invalid: %s", subject)
	}
	if len(message) <= 0 {
		return fmt.Errorf("message is invalid: %s", message)
	}
	if err := e.preReq(); err != nil {
		return err
	}

	m := mail.NewV3Mail()
	from := mail.NewEmail(e.senderName, e.senderEmail)
	m.SetFrom(from)
	m.Subject = subject
	m.AddContent(mail.NewContent("text/plain", message))

	tos := []*mail.Email{}
	for _, _to := range to {
		toName := _to
		toEmail := _to
		tos = append(tos, mail.NewEmail(toName, toEmail))
	}
	p := mail.NewPersonalization()
	p.AddTos(tos...)
	m.AddPersonalizations(p)

	trackingSettings := mail.NewTrackingSettings()

	clickTrackingSetting := mail.NewClickTrackingSetting()
	clickTrackingSetting.SetEnable(false)
	clickTrackingSetting.SetEnableText(false)
	trackingSettings.SetClickTracking(clickTrackingSetting)

	openTrackingSetting := mail.NewOpenTrackingSetting()
	openTrackingSetting.SetEnable(false)
	trackingSettings.SetOpenTracking(openTrackingSetting)

	subscriptionTrackingSetting := mail.NewSubscriptionTrackingSetting()
	subscriptionTrackingSetting.SetEnable(false)
	trackingSettings.SetSubscriptionTracking(subscriptionTrackingSetting)

	m.SetTrackingSettings(trackingSettings)

	request := sendgrid.GetRequest(e.sendGridApiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(m)
	request.Body = Body
	_, err := sendgrid.API(request)
	if err != nil {
		return err
	}

	logger.With("email", "sendgrid").Debug("Email sent", "to", to, "subject", subject)
	return nil
}

func (e *emailSendgrid) Options() EmailOptions {
	return e.options
}

func (e *emailSendgrid) String() string {
	return "EmailSendgrid"
}

func (e *emailSendgrid) preReq() error {
	if len(e.sendGridApiKey) <= 0 {
		return fmt.Errorf("sendgrid api key is not set")
	}
	if len(e.senderEmail) <= 0 {
		return fmt.Errorf("senderEmail is not set")
	}
	if len(e.senderName) <= 0 {
		return fmt.Errorf("name is not set")
	}
	return nil
}
