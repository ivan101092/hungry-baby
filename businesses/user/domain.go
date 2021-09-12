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
	Credentials     []DomainCredentials
	Status          bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type DomainSettings struct {
	AutoNotification bool `json:"auto_notification"`
}

type DomainCredentials struct {
	ID                  int                       `json:"id"`
	Type                string                    `json:"type"`
	Email               string                    `json:"email"`
	RegistrationDetails DomainRegistrationDetails `json:"registration_details"`
}

type DomainRegistrationDetails struct {
	AccessToken  string `json:"access_token"`
	Expiry       string `json:"expiry"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
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
