package response

import (
	"hungry-baby/businesses/userChild"
	"time"
)

type UserChild struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	Name              string    `json:"name"`
	Gender            string    `json:"gender"`
	BirthDate         string    `json:"birth_date"`
	BirthLength       float64   `json:"birth_length"`
	BirthWeight       float64   `json:"birth_weight"`
	HeadCircumference float64   `json:"head_circumference"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func FromDomain(domain userChild.Domain) UserChild {
	return UserChild{
		ID:                domain.ID,
		UserID:            domain.UserID,
		Name:              domain.Name,
		Gender:            domain.Gender,
		BirthDate:         domain.BirthDate,
		BirthLength:       domain.BirthLength,
		BirthWeight:       domain.BirthWeight,
		HeadCircumference: domain.HeadCircumference,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}
