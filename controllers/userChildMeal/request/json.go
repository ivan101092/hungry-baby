package request

import (
	"hungry-baby/businesses/userChildMeal"
)

type UserChildMeal struct {
	UserID      int
	UserChildID int     `json:"user_child_id"`
	MealPlanID  int     `json:"meal_plan_id"`
	Quantity    float64 `json:"quantity"`
	ScheduledAt string  `json:"schedule_at"`
	FinishAt    string  `json:"finish_at"`
}

func (req *UserChildMeal) ToDomain() *userChildMeal.Domain {
	return &userChildMeal.Domain{
		UserID:      req.UserID,
		UserChildID: req.UserChildID,
		MealPlanID:  req.MealPlanID,
		Quantity:    req.Quantity,
		ScheduledAt: req.ScheduledAt,
		FinishAt:    req.FinishAt,
	}
}
