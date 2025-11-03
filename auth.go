package auth

import (
	"os"

	"github.com/gradientzero/comby/v2"
	"github.com/joho/godotenv"
)

var (
	// ACCOUNT_REGISTRATION_CONFIRMATION_URL is the default URL for account registration confirmation.
	// While developing, this URL can be set to a local URL like "http://localhost:9000/auth/register/{ott}".
	ACCOUNT_REGISTRATION_CONFIRMATION_URL string = "https://www.example.tld/auth/register/{ott}"

	// account account confirmation email content
	ACCOUNT_REGISTRATION_CONFIRMATION_SUBJECT string = "Confirm your registration"
	ACCOUNT_REGISTRATION_CONFIRMATION_MESSAGE string = "Please confirm your registration by clicking the following link: {linkUrl}"
)

// init initializes system-relevant variables from environment variables if provided.
func init() {
	// ensure to load local .env file
	godotenv.Load()
	{
		// Account Registration Confirmation URL
		providedValue := os.Getenv("ACCOUNT_REGISTRATION_CONFIRMATION_URL")
		if len(providedValue) > 0 {
			ACCOUNT_REGISTRATION_CONFIRMATION_URL = providedValue
		}
	}
	{
		// Account Registration Confirmation Email Subject
		providedValue := os.Getenv("ACCOUNT_REGISTRATION_CONFIRMATION_SUBJECT")
		if len(providedValue) > 0 {
			ACCOUNT_REGISTRATION_CONFIRMATION_SUBJECT = providedValue
		}
	}
	{
		// Account Registration Confirmation Email Message
		providedValue := os.Getenv("ACCOUNT_REGISTRATION_CONFIRMATION_MESSAGE")
		if len(providedValue) > 0 {
			ACCOUNT_REGISTRATION_CONFIRMATION_MESSAGE = providedValue
		}
	}
}

var logger = comby.Logger.With("pkg", "comby/auth")
