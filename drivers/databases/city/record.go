package city

import (
	"hungry-baby/businesses/city"
	cityUsecase "hungry-baby/businesses/city"
	"hungry-baby/drivers/postgres"
)

type City struct {
	ID         int
	ProvinceID int
	Code       string
	Name       string
	Status     bool
	postgres.BaseModel
}

func FromDomain(domain *cityUsecase.Domain) *City {
	return &City{
		ID:         domain.ID,
		ProvinceID: domain.ProvinceID,
		Code:       domain.Code,
		Name:       domain.Name,
		Status:     domain.Status,
	}
}

func (rec *City) ToDomain() city.Domain {
	return city.Domain{
		ID:         rec.ID,
		ProvinceID: rec.ProvinceID,
		Code:       rec.Code,
		Name:       rec.Name,
		Status:     rec.Status,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}
