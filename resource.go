package auth

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gradientzero/comby-auth-email/email"
	"github.com/gradientzero/comby/v2"
)

// Resource struct
type Resource struct {
	fc    *comby.Facade
	api   huma.API
	email email.Email
}

// NewResource creates new instance
func NewResource(fc *comby.Facade, api huma.API, email email.Email) *Resource {
	resource := &Resource{
		fc:    fc,
		api:   api,
		email: email,
	}
	return resource
}

// AuthScope represents the authorization scope for a specific action.
type AuthScope struct {
	OneTimeToken string `json:"oneTimeToken,omitempty"`
	Action       string `json:"action"`
	Object       string `json:"object"`
	ExpiredAt    int64  `json:"expiredAt"`
}

type ResponseAuthWithScope struct {
	Body struct {
		Scope   *AuthScope `json:"authScope,omitempty"`
		Error   string     `json:"error,omitempty"`
		Message string     `json:"message,omitempty"`
	}
}

func (rs *Resource) Register() {
	huma.Register(rs.api, huma.Operation{
		OperationID: "auth-account-register",
		Description: "Register new account",
		Method:      http.MethodPost,
		Path:        "/api/auth/register/:ott",
		Tags:        []string{"Auth"},
	}, rs.AccountRegister)

	huma.Register(rs.api, huma.Operation{
		OperationID: "auth-ott-validate",
		Description: "Validate one-time token",
		Method:      http.MethodGet,
		Path:        "/api/auth/ott/:ott/validate",
		Tags:        []string{"Auth"},
	}, rs.OneTimeTokenValidate)
}
