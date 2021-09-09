package country

import (
	"hungry-baby/businesses/country"
	countryUsecase "hungry-baby/businesses/country"
	"hungry-baby/drivers/postgres"
)

type Country struct {
	ID          int
	CountryCode string
	Name        string
	Status      bool
	postgres.BaseModel
}

func FromDomain(domain *countryUsecase.Domain) *Country {
	return &Country{
		ID:          domain.ID,
		CountryCode: domain.CountryCode,
		Name:        domain.Name,
		Status:      domain.Status,
	}
}

func (rec *Country) ToDomain() country.Domain {
	return country.Domain{
		ID:          rec.ID,
		CountryCode: rec.CountryCode,
		Name:        rec.Name,
		Status:      rec.Status,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}
