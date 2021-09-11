package city

import (
	"context"
	"time"
)

type Domain struct {
	ID         int
	ProvinceID int
	Code       string
	Name       string
	Status     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	FindAll(ctx context.Context, search string, provinceID int, status string) ([]Domain, error)
	Find(ctx context.Context, search string, provinceID int, status string, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int, status string) (Domain, error)
	FindByCode(ctx context.Context, code, status string) (Domain, error)
	Store(ctx context.Context, cityDomain *Domain) (Domain, error)
	Update(ctx context.Context, cityDomain *Domain) (Domain, error)
	Delete(ctx context.Context, cityDomain *Domain) (Domain, error)
}

type Repository interface {
	FindAll(ctx context.Context, search string, provinceID int, status string) ([]Domain, error)
	Find(ctx context.Context, search string, provinceID int, status string, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int, status string) (Domain, error)
	FindByCode(ctx context.Context, code, status string) (Domain, error)
	Store(ctx context.Context, cityDomain *Domain) (Domain, error)
	Update(ctx context.Context, cityDomain *Domain) (Domain, error)
	Delete(ctx context.Context, cityDomain *Domain) (Domain, error)
}
