package request

import (
	"hungry-baby/businesses/userChild"
)

type UserChild struct {
	UserID            int     `json:"user_id"`
	Name              string  `json:"name"`
	Gender            string  `json:"gender"`
	BirthDate         string  `json:"birth_date"`
	BirthLength       float64 `json:"birth_length"`
	BirthWeight       float64 `json:"birth_weight"`
	HeadCircumference float64 `json:"head_circumference"`
}

func (req *UserChild) ToDomain() *userChild.Domain {
	return &userChild.Domain{
		UserID:            req.UserID,
		Name:              req.Name,
		Gender:            req.Gender,
		BirthDate:         req.BirthDate,
		BirthLength:       req.BirthLength,
		BirthWeight:       req.BirthWeight,
		HeadCircumference: req.HeadCircumference,
	}
}
