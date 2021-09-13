package userChildMeal

import (
	"hungry-baby/businesses/userChildMeal"
	userChildMealUsecase "hungry-baby/businesses/userChildMeal"
	"hungry-baby/drivers/postgres"
	"time"
)

type UserChildMeal struct {
	ID                 int
	UserChildID        int
	MealPlanID         int
	Name               string
	SuggestionQuantity float64
	Quantity           float64
	Unit               string
	ScheduledAt        time.Time
	FinishAt           time.Time
	CalendarID         string
	Status             string
	postgres.BaseModel
}

func FromDomain(domain *userChildMealUsecase.Domain) *UserChildMeal {
	scheduleAt, _ := time.Parse(time.RFC3339, domain.ScheduledAt)
	finishAt, _ := time.Parse(time.RFC3339, domain.FinishAt)

	return &UserChildMeal{
		ID:                 domain.ID,
		UserChildID:        domain.UserChildID,
		MealPlanID:         domain.MealPlanID,
		Name:               domain.Name,
		SuggestionQuantity: domain.SuggestionQuantity,
		Quantity:           domain.Quantity,
		Unit:               domain.Unit,
		ScheduledAt:        scheduleAt,
		FinishAt:           finishAt,
		CalendarID:         domain.CalendarID,
		Status:             domain.Status,
	}
}

func (rec *UserChildMeal) ToDomain() userChildMeal.Domain {
	return userChildMeal.Domain{
		ID:                 rec.ID,
		UserChildID:        rec.UserChildID,
		MealPlanID:         rec.MealPlanID,
		Name:               rec.Name,
		SuggestionQuantity: rec.SuggestionQuantity,
		Quantity:           rec.Quantity,
		Unit:               rec.Unit,
		ScheduledAt:        rec.ScheduledAt.Format(time.RFC3339),
		FinishAt:           rec.FinishAt.Format(time.RFC3339),
		CalendarID:         rec.CalendarID,
		Status:             rec.Status,
		CreatedAt:          rec.CreatedAt,
		UpdatedAt:          rec.UpdatedAt,
	}
}
