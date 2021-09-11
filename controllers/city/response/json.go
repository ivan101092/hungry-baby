package response

import (
	"hungry-baby/businesses/city"
	"time"
)

type City struct {
	ID         int       `json:"id"`
	ProvinceID int       `json:"province_id"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomain(domain city.Domain) City {
	return City{
		ID:         domain.ID,
		ProvinceID: domain.ProvinceID,
		Code:       domain.Code,
		Name:       domain.Name,
		Status:     domain.Status,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
