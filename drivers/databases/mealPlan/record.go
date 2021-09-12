package mealPlan

import (
	"hungry-baby/businesses/mealPlan"
	mealPlanUsecase "hungry-baby/businesses/mealPlan"
	"hungry-baby/drivers/postgres"
)

type MealPlan struct {
	ID                 int
	UserID             int
	Name               string
	MinAge             float64
	MaxAge             float64
	Interval           float64
	SuggestionQuantity float64
	Unit               string
	Status             bool
	postgres.BaseModel
}

func FromDomain(domain *mealPlanUsecase.Domain) *MealPlan {
	return &MealPlan{
		ID:                 domain.ID,
		UserID:             domain.UserID,
		Name:               domain.Name,
		MinAge:             domain.MinAge,
		MaxAge:             domain.MaxAge,
		Interval:           domain.Interval,
		SuggestionQuantity: domain.SuggestionQuantity,
		Unit:               domain.Unit,
		Status:             domain.Status,
	}
}

func (rec *MealPlan) ToDomain() mealPlan.Domain {
	return mealPlan.Domain{
		ID:                 rec.ID,
		UserID:             rec.UserID,
		Name:               rec.Name,
		MinAge:             rec.MinAge,
		MaxAge:             rec.MaxAge,
		Interval:           rec.Interval,
		SuggestionQuantity: rec.SuggestionQuantity,
		Unit:               rec.Unit,
		Status:             rec.Status,
		CreatedAt:          rec.CreatedAt,
		UpdatedAt:          rec.UpdatedAt,
	}
}
