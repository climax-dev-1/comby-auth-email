package auth

import (
	"context"
)

type RequestOneTimeTokenValidate struct {
	Body struct {
		OneTimeToken string `json:"oneTimeToken" validate:"required"`
	}
}

func (rs *Resource) OneTimeTokenValidate(ctx context.Context, req *RequestOneTimeTokenValidate) (*ResponseAuthWithScope, error) {
	// ...

	// create response
	resp := &ResponseAuthWithScope{}
	return resp, nil
}
