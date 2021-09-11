package response

import (
	"hungry-baby/businesses/province"
	"time"
)

type Province struct {
	ID        int       `json:"id"`
	CountryID int       `json:"country_id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain province.Domain) Province {
	return Province{
		ID:        domain.ID,
		CountryID: domain.CountryID,
		Code:      domain.Code,
		Name:      domain.Name,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
