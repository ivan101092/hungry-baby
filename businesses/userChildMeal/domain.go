package userChildMeal

import (
	"context"
	"time"
)

type Domain struct {
	ID                 int
	UserID             int
	UserChildID        int
	MealPlanID         int
	Name               string
	SuggestionQuantity float64
	Quantity           float64
	Unit               string
	ScheduledAt        string
	FinishAt           string
	CalendarID         string
	Status             string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Usecase interface {
	FindAll(ctx context.Context, search string, userChildID int) ([]Domain, error)
	Find(ctx context.Context, search string, userChildID, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	FindByChildMeal(ctx context.Context, userChildID, mealPlanID int) (Domain, error)
	Store(ctx context.Context, userChildMealDomain *Domain) (Domain, error)
	Update(ctx context.Context, userChildMealDomain *Domain) (Domain, error)
	Delete(ctx context.Context, userChildMealDomain *Domain) (Domain, error)
}

type Repository interface {
	FindAll(ctx context.Context, search string, userChildID int) ([]Domain, error)
	Find(ctx context.Context, search string, userChildID, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	FindByChildMeal(ctx context.Context, userChildID, mealPlanID int) (Domain, error)
	FindNextPending(ctx context.Context, userChildID, mealPlanID int) (Domain, error)
	Store(ctx context.Context, userChildMealDomain *Domain) (Domain, error)
	Update(ctx context.Context, userChildMealDomain *Domain) (Domain, error)
	Delete(ctx context.Context, userChildMealDomain *Domain) (Domain, error)
}
