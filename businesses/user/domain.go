package user

import (
	"context"
	"time"
)

type Domain struct {
	ID              int
	Code            string
	Email           string
	Name            string
	ProfileImageID  int
	ProfileImageURL string
	Gender          string
	Phone           string
	CityID          int
	CityName        string
	Address         string
	Settings        DomainSettings
	Status          bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type DomainSettings struct {
	AutoNotification bool `json:"auto_notification"`
}

type Usecase interface {
	FindAll(ctx context.Context, search string, status string) ([]Domain, error)
	Find(ctx context.Context, search string, status string, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int, status string) (Domain, error)
	FindByCode(ctx context.Context, code, status string) (Domain, error)
	FindByEmail(ctx context.Context, email, status string) (Domain, error)
	Store(ctx context.Context, userDomain *Domain) (Domain, error)
	Update(ctx context.Context, userDomain *Domain) (Domain, error)
	Delete(ctx context.Context, userDomain *Domain) (Domain, error)
}

type Repository interface {
	FindAll(ctx context.Context, search string, status string) ([]Domain, error)
	Find(ctx context.Context, search string, status string, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int, status string) (Domain, error)
	FindByCode(ctx context.Context, code, status string) (Domain, error)
	FindByEmail(ctx context.Context, email, status string) (Domain, error)
	Store(ctx context.Context, userDomain *Domain) (Domain, error)
	Update(ctx context.Context, userDomain *Domain) (Domain, error)
	Delete(ctx context.Context, userDomain *Domain) (Domain, error)
}
