package response

import (
	"hungry-baby/businesses/user"
	"hungry-baby/helpers/interfacepkg"
)

type User struct {
	ID              int               `json:"id"`
	Code            string            `json:"code"`
	Email           string            `json:"email"`
	Name            string            `json:"name"`
	ProfileImageID  int               `json:"profile_image_id"`
	ProfileImageURL string            `json:"profile_image_url"`
	Gender          string            `json:"gender"`
	Phone           string            `json:"phone"`
	CityID          int               `json:"city_id"`
	CityName        string            `json:"city_name"`
	Address         string            `json:"address"`
	Settings        UserSettings      `json:"settings"`
	Credentials     []UserCredentials `json:"credentials"`
	Status          bool              `json:"status"`
}

type UserSettings struct {
	AutoNotification bool `json:"auto_notification"`
}

type UserCredentials struct {
	ID                  int                     `json:"id"`
	Type                string                  `json:"type"`
	Email               string                  `json:"email"`
	RegistrationDetails UserRegistrationDetails `json:"registration_details"`
}

type UserRegistrationDetails struct {
	AccessToken  string `json:"access_token"`
	Expiry       string `json:"expiry"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

func FromDomain(domain user.Domain) User {
	var settings UserSettings
	interfacepkg.UnmarshalCbInterface(domain.Settings, &settings)

	var credentials []UserCredentials
	interfacepkg.UnmarshalCbInterface(domain.Credentials, &credentials)

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
		Credentials:     credentials,
		Status:          domain.Status,
	}
}
