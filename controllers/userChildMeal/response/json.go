package response

import (
	"hungry-baby/businesses/userChildMeal"
	"time"
)

type UserChildMeal struct {
	ID                 int       `json:"id"`
	UserChildID        int       `json:"user_child_id"`
	MealPlanID         int       `json:"meal_plan_id"`
	Name               string    `json:"name"`
	SuggestionQuantity float64   `json:"suggestion_quantity"`
	Quantity           float64   `json:"quantity"`
	Unit               string    `json:"unit"`
	ScheduledAt        string    `json:"schedule_at"`
	FinishAt           string    `json:"finish_at"`
	CalendarID         string    `json:"calendar_id"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func FromDomain(domain userChildMeal.Domain) UserChildMeal {
	return UserChildMeal{
		ID:                 domain.ID,
		UserChildID:        domain.UserChildID,
		MealPlanID:         domain.MealPlanID,
		Name:               domain.Name,
		SuggestionQuantity: domain.SuggestionQuantity,
		Quantity:           domain.Quantity,
		Unit:               domain.Unit,
		ScheduledAt:        domain.ScheduledAt,
		FinishAt:           domain.FinishAt,
		CalendarID:         domain.CalendarID,
		Status:             domain.Status,
		CreatedAt:          domain.CreatedAt,
		UpdatedAt:          domain.UpdatedAt,
	}
}
