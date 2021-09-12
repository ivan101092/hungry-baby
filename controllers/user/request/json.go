package request

import (
	"hungry-baby/businesses/user"
	"hungry-baby/helpers/interfacepkg"
)

type User struct {
	ProvinceID     int          `json:"province_id"`
	Code           string       `json:"code"`
	Email          string       `json:"email"`
	Name           string       `json:"name"`
	ProfileImageID int          `json:"profile_image_id"`
	Gender         string       `json:"gender"`
	Phone          string       `json:"phone"`
	CityID         int          `json:"city_id"`
	Address        string       `json:"address"`
	Settings       UserSettings `json:"settings"`
	Status         bool         `json:"status"`
}

type UserSettings struct {
	AutoNotification bool `json:"auto_notification"`
}

func (req *User) ToDomain() *user.Domain {
	var settings user.DomainSettings
	interfacepkg.UnmarshalCbInterface(req.Settings, &settings)

	return &user.Domain{
		Code:           req.Code,
		Email:          req.Email,
		Name:           req.Name,
		ProfileImageID: req.ProfileImageID,
		Gender:         req.Gender,
		Phone:          req.Phone,
		CityID:         req.CityID,
		Address:        req.Address,
		Settings:       settings,
		Status:         req.Status,
	}
}
