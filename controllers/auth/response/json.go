package response

import (
	"hungry-baby/businesses/auth"
)

type Auth struct {
	Token     string `json:"token"`
	ExpiredAt int    `json:"expired_at"`
}

func FromDomain(domain auth.Domain) Auth {
	return Auth{
		Token:     domain.Token,
		ExpiredAt: domain.ExpiredAt,
	}
}
