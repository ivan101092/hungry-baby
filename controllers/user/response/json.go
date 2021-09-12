package response

import (
	"hungry-baby/businesses/user"
	"hungry-baby/helpers/interfacepkg"
)

type User struct {
	ID              int          `json:"id"`
	Code            string       `json:"code"`
	Email           string       `json:"email"`
	Name            string       `json:"name"`
	ProfileImageID  int          `json:"profile_image_id"`
	ProfileImageURL string       `json:"profile_image_url"`
	Gender          string       `json:"gender"`
	Phone           string       `json:"phone"`
	CityID          int          `json:"city_id"`
	CityName        string       `json:"city_name"`
	Address         string       `json:"address"`
	Settings        UserSettings `json:"settings"`
	Status          bool         `json:"status"`
}

type UserSettings struct {
	AutoNotification bool `json:"auto_notification"`
}

func FromDomain(domain user.Domain) User {
	var settings UserSettings
	interfacepkg.UnmarshalCbInterface(domain.Settings, &settings)

	return User{
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
		Settings:        settings,
		Status:          domain.Status,
	}
}
