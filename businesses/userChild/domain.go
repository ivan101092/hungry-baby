package userChild

import (
	"context"
	"time"
)

type Domain struct {
	ID                int
	UserID            int
	Name              string
	Gender            string
	BirthDate         string
	BirthLength       float64
	BirthWeight       float64
	HeadCircumference float64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Usecase interface {
	FindAll(ctx context.Context, search string, userID int) ([]Domain, error)
	Find(ctx context.Context, search string, userID, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, userChildDomain *Domain) (Domain, error)
	Update(ctx context.Context, userChildDomain *Domain) (Domain, error)
	Delete(ctx context.Context, userChildDomain *Domain) (Domain, error)
}

type Repository interface {
	FindAll(ctx context.Context, search string, userID int) ([]Domain, error)
	Find(ctx context.Context, search string, userID, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, userChildDomain *Domain) (Domain, error)
	Update(ctx context.Context, userChildDomain *Domain) (Domain, error)
	Delete(ctx context.Context, userChildDomain *Domain) (Domain, error)
}
