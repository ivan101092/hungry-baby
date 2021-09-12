package mealPlan

import (
	"context"
	"time"
)

type Domain struct {
	ID                 int
	UserID             int
	Name               string
	MinAge             float64
	MaxAge             float64
	Interval           float64
	SuggestionQuantity float64
	Unit               string
	Status             bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Usecase interface {
	FindAll(ctx context.Context, search string, userID int, status string) ([]Domain, error)
	Find(ctx context.Context, search string, userID int, status string, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int, status string) (Domain, error)
	Store(ctx context.Context, mealPlanDomain *Domain) (Domain, error)
	Update(ctx context.Context, mealPlanDomain *Domain) (Domain, error)
	Delete(ctx context.Context, mealPlanDomain *Domain) (Domain, error)
}

type Repository interface {
	FindAll(ctx context.Context, search string, userID int, status string) ([]Domain, error)
	Find(ctx context.Context, search string, userID int, status string, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int, status string) (Domain, error)
	Store(ctx context.Context, mealPlanDomain *Domain) (Domain, error)
	Update(ctx context.Context, mealPlanDomain *Domain) (Domain, error)
	Delete(ctx context.Context, mealPlanDomain *Domain) (Domain, error)
}
