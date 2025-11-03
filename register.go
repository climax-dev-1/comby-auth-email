package auth

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gradientzero/comby/v2"
	"github.com/gradientzero/comby/v2/api"
)

const AUTH_SCOPE_REGISTER = "auth-account-register"

type RequestAuthAccountRegister struct {
	Body struct {
		Email string `json:"email" validate:"required" doc:"email of the account" example:"user@example.com"`
	}
}

func (rs *Resource) AccountRegister(ctx context.Context, req *RequestAuthAccountRegister) (*ResponseAuthWithScope, error) {

	// validate
	if err := comby.ValidateEmail(req.Body.Email); err != nil {
		return nil, api.SchemaError(err)
	}

	// filter and throttle requests
	if err := rs.ThrottleRequest(ctx, req); err != nil {
		return nil, api.SchemaError(err)
	}

	// create scope
	expirationDuration := 15 * time.Minute
	scope := &AuthScope{
		OneTimeToken: comby.RandomDigits(6),
		Action:       AUTH_SCOPE_REGISTER,
		Object:       req.Body.Email,
		ExpiredAt:    time.Now().Add(expirationDuration).UnixNano(),
	}
	scopeBytes, err := comby.Serialize(scope)
	if err != nil {
		return nil, api.SchemaError(err)
	}

	// store scope in CacheStore
	scopeKey := fmt.Sprintf("auth-scope-%s-%s", req.Body.Email, scope.OneTimeToken)
	if err := rs.fc.GetCacheStore().Set(ctx,
		comby.CacheStoreSetOptionWithKeyValue(scopeKey, scopeBytes),
		comby.CacheStoreSetOptionWithExpiration(expirationDuration),
	); err != nil {
		return nil, api.SchemaError(err)
	}

	// send email with one-time token
	uri := ACCOUNT_REGISTRATION_CONFIRMATION_URL
	emailTo := req.Body.Email
	linkUrl := fmt.Sprintf("%s/%s", uri, scope.OneTimeToken)
	subject := ACCOUNT_REGISTRATION_CONFIRMATION_SUBJECT
	message := ACCOUNT_REGISTRATION_CONFIRMATION_MESSAGE
	message = strings.ReplaceAll(message, "{linkUrl}", linkUrl)
	if rs.email != nil {
		if err := rs.email.SendMail([]string{emailTo}, subject, message); err != nil {
			logger.Error("failed to send account registration confirmation email",
				"to", emailTo, "subject", subject,
				"link", linkUrl,
				"error", err)
			return nil, api.SchemaError(err)
		}
	} else {
		logger.Warn("Email service not configured, skipping sending confirmation email",
			"to", emailTo, "subject", subject,
			"link", linkUrl)
	}

	// create response
	resp := &ResponseAuthWithScope{}
	resp.Body.Scope = scope
	return resp, nil
}
