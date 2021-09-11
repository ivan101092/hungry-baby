package province

import (
	"hungry-baby/businesses/province"
	provinceUsecase "hungry-baby/businesses/province"
	"hungry-baby/drivers/postgres"
)

type Province struct {
	ID        int
	CountryID int
	Code      string
	Name      string
	Status    bool
	postgres.BaseModel
}

func FromDomain(domain *provinceUsecase.Domain) *Province {
	return &Province{
		ID:        domain.ID,
		CountryID: domain.CountryID,
		Code:      domain.Code,
		Name:      domain.Name,
		Status:    domain.Status,
	}
}

func (rec *Province) ToDomain() province.Domain {
	return province.Domain{
		ID:        rec.ID,
		CountryID: rec.CountryID,
		Code:      rec.Code,
		Name:      rec.Name,
		Status:    rec.Status,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
