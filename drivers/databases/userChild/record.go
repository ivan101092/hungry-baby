package userChild

import (
	"hungry-baby/businesses/userChild"
	userChildUsecase "hungry-baby/businesses/userChild"
	"hungry-baby/drivers/postgres"
)

type UserChild struct {
	ID                int
	UserID            int
	Name              string
	Gender            string
	BirthDate         string
	BirthLength       float64
	BirthWeight       float64
	HeadCircumference float64
	postgres.BaseModel
}

func FromDomain(domain *userChildUsecase.Domain) *UserChild {
	return &UserChild{
		ID:                domain.ID,
		UserID:            domain.UserID,
		Name:              domain.Name,
		Gender:            domain.Gender,
		BirthDate:         domain.BirthDate,
		BirthLength:       domain.BirthLength,
		BirthWeight:       domain.BirthWeight,
		HeadCircumference: domain.HeadCircumference,
	}
}

func (rec *UserChild) ToDomain() userChild.Domain {
	return userChild.Domain{
		ID:                rec.ID,
		UserID:            rec.UserID,
		Name:              rec.Name,
		Gender:            rec.Gender,
		BirthDate:         rec.BirthDate,
		BirthLength:       rec.BirthLength,
		BirthWeight:       rec.BirthWeight,
		HeadCircumference: rec.HeadCircumference,
		CreatedAt:         rec.CreatedAt,
		UpdatedAt:         rec.UpdatedAt,
	}
}
