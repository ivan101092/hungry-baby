package response

import (
	"hungry-baby/businesses/country"
	"time"
)

type Country struct {
	ID          int       `json:"id"`
	CountryCode string    `json:"country_code"`
	Name        string    `json:"name"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain country.Domain) Country {
	return Country{
		ID:          domain.ID,
		CountryCode: domain.CountryCode,
		Name:        domain.Name,
		Status:      domain.Status,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
