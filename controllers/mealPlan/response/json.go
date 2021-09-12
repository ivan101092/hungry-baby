package response

import (
	"hungry-baby/businesses/mealPlan"
	"time"
)

type MealPlan struct {
	ID                 int       `json:"id"`
	UserID             int       `json:"user_id"`
	Name               string    `json:"name"`
	MinAge             float64   `json:"min_age"`
	MaxAge             float64   `json:"max_age"`
	Interval           float64   `json:"interval"`
	SuggestionQuantity float64   `json:"suggestion_quantity"`
	Unit               string    `json:"unit"`
	Status             bool      `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func FromDomain(domain mealPlan.Domain) MealPlan {
	return MealPlan{
		ID:                 domain.ID,
		UserID:             domain.UserID,
		Name:               domain.Name,
		MinAge:             domain.MinAge,
		MaxAge:             domain.MaxAge,
		Interval:           domain.Interval,
		SuggestionQuantity: domain.SuggestionQuantity,
		Unit:               domain.Unit,
		Status:             domain.Status,
		CreatedAt:          domain.CreatedAt,
		UpdatedAt:          domain.UpdatedAt,
	}
}
