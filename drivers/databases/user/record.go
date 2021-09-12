package user

import (
	"hungry-baby/businesses/user"
	userUsecase "hungry-baby/businesses/user"
	"hungry-baby/drivers/postgres"
	"hungry-baby/helpers/interfacepkg"
)

type User struct {
	ID              int
	Code            string
	Email           string
	Name            string
	ProfileImageID  int
	ProfileImageURL string `json:"profile_image_url" gorm:"<-:false"`
	Gender          string
	Phone           string
	CityID          int
	CityName        string `json:"city_name" gorm:"<-:false"`
	Address         string
	Settings        string
	Status          bool
	postgres.BaseModel
}

func FromDomain(domain *userUsecase.Domain) *User {
	return &User{
		ID:              domain.ID,
		Code:            domain.Code,
		Email:           domain.Email,
		Name:            domain.Name,
		ProfileImageID:  domain.ProfileImageID,
		ProfileImageURL: domain.ProfileImageURL,
		Gender:          domain.Gender,
		Phone:           domain.Phone,
		CityID:          domain.CityID,
		CityName:        domain.CityName,
		Address:         domain.Address,
		Settings:        interfacepkg.Marshal(domain.Settings),
		Status:          domain.Status,
	}
}

func (rec *User) ToDomain() user.Domain {
	var settings user.DomainSettings
	interfacepkg.UnmarshalCb(rec.Settings, &settings)

	return user.Domain{
		ID:              rec.ID,
		Code:            rec.Code,
		Email:           rec.Email,
		Name:            rec.Name,
		ProfileImageID:  rec.ProfileImageID,
		ProfileImageURL: rec.ProfileImageURL,
		Gender:          rec.Gender,
		Phone:           rec.Phone,
		CityID:          rec.CityID,
		CityName:        rec.CityName,
		Address:         rec.Address,
		Settings:        settings,
		Status:          rec.Status,
	}
}
