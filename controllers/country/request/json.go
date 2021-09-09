package request

import (
	"hungry-baby/businesses/country"
)

type Country struct {
	CountryCode string `json:"country_code"`
	Name        string `json:"name"`
	Status      bool   `json:"status"`
}

func (req *Country) ToDomain() *country.Domain {
	return &country.Domain{
		CountryCode: req.CountryCode,
		Name:        req.Name,
		Status:      req.Status,
	}
}
