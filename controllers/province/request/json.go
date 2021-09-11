package request

import (
	"hungry-baby/businesses/province"
)

type Province struct {
	CountryID int    `json:"country_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Status    bool   `json:"status"`
}

func (req *Province) ToDomain() *province.Domain {
	return &province.Domain{
		CountryID: req.CountryID,
		Code:      req.Code,
		Name:      req.Name,
		Status:    req.Status,
	}
}
