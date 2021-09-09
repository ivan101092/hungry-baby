package country

import (
	"context"
	"time"
)

type Domain struct {
	ID          int
	CountryCode string
	Name        string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	FindAll(ctx context.Context, search, status string) ([]Domain, error)
	Find(ctx context.Context, search, status string, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int, status string) (Domain, error)
	FindByCode(ctx context.Context, code, status string) (Domain, error)
	Store(ctx context.Context, countryDomain *Domain) (Domain, error)
	Update(ctx context.Context, countryDomain *Domain) (Domain, error)
	Delete(ctx context.Context, countryDomain *Domain) (Domain, error)
}

type Repository interface {
	FindAll(ctx context.Context, search, status string) ([]Domain, error)
	Find(ctx context.Context, search, status string, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int, status string) (Domain, error)
	FindByCode(ctx context.Context, code, status string) (Domain, error)
	Store(ctx context.Context, countryDomain *Domain) (Domain, error)
	Update(ctx context.Context, countryDomain *Domain) (Domain, error)
	Delete(ctx context.Context, countryDomain *Domain) (Domain, error)
}
