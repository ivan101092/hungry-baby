package request

import (
	"hungry-baby/businesses/mealPlan"
)

type MealPlan struct {
	UserID             int     `json:"user_id"`
	Name               string  `json:"name"`
	MinAge             float64 `json:"min_age"`
	MaxAge             float64 `json:"max_age"`
	Interval           float64 `json:"interval"`
	SuggestionQuantity float64 `json:"suggestion_quantity"`
	Unit               string  `json:"unit"`
	Status             bool    `json:"status"`
}

func (req *MealPlan) ToDomain() *mealPlan.Domain {
	return &mealPlan.Domain{
		UserID:             req.UserID,
		Name:               req.Name,
		MinAge:             req.MinAge,
		MaxAge:             req.MaxAge,
		Interval:           req.Interval,
		SuggestionQuantity: req.SuggestionQuantity,
		Unit:               req.Unit,
		Status:             req.Status,
	}
}
