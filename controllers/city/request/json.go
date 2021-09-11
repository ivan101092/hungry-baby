package request

import (
	"hungry-baby/businesses/city"
)

type City struct {
	ProvinceID int    `json:"province_id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Status     bool   `json:"status"`
}

func (req *City) ToDomain() *city.Domain {
	return &city.Domain{
		ProvinceID: req.ProvinceID,
		Code:       req.Code,
		Name:       req.Name,
		Status:     req.Status,
	}
}
