package auth

import "context"

type Domain struct {
	Token     string `json:"token"`
	ExpiredAt int    `json:"expired_at"`
}

type Repository interface {
}

type Usecase interface {
	GetGoogleLoginURL(ctx context.Context) (string, error)
	VerifyGoogleCode(ctx context.Context, code string) (Domain, error)
}
