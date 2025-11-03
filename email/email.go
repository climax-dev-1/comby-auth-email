package email

import "github.com/gradientzero/comby/v2"

// EmailOptions represents configuration options for the Email service.
//
// It is currently an empty struct, but can be extended in the future.
type EmailOptions struct {
	// No fields defined at the moment.
}

// EmailOption defines a function type for setting options on the Email service.
//
// It allows for configuring the Email service by applying various options.
type EmailOption func(opt *EmailOptions) (*EmailOptions, error)

// Email defines an interface for an email service.
//
// It provides methods for initializing the service, sending general emails, and sending specific types of emails
// such as account password reset and invitations.
type Email interface {
	// Init initializes the email service with the provided options.
	Init(opts ...EmailOption) error

	// SendMail sends an email to the specified recipients.
	//
	// Parameters:
	// - to: List of email addresses to send the email to.
	// - subject: Subject of the email.
	// - message: Body of the email.
	SendMail(to []string, subject, message string) error

	// Options returns the current configuration options of the email service.
	Options() EmailOptions

	// String returns a string representation of the email service.
	String() string
}

var logger = comby.Logger.With("pkg", "comby/auth/email")
