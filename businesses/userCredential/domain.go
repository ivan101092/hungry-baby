package userCredential

import (
	"context"
	"time"
)

type Domain struct {
	ID                  int
	UserID              int
	Type                string
	Email               string
	EmailValidAt        string
	Password            string
	RegistrationDetails DomainRegistrationDetails
	Status              bool
	CreatedAt           time.Time
	UpdatedAt           time.Time
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
	FindByEmail(ctx context.Context, email, status string) (Domain, error)
	FindByUserType(ctx context.Context, userID int, types, status string) (Domain, error)
	Store(ctx context.Context, userDomain *Domain) (Domain, error)
	Update(ctx context.Context, userDomain *Domain) (Domain, error)
	Delete(ctx context.Context, userDomain *Domain) (Domain, error)
}

type Repository interface {
	FindAll(ctx context.Context, search string, status string) ([]Domain, error)
	Find(ctx context.Context, search string, status string, page, perpage int) ([]Domain, int, error)
	FindByID(ctx context.Context, id int, status string) (Domain, error)
	FindByEmail(ctx context.Context, email, status string) (Domain, error)
	FindByUserType(ctx context.Context, userID int, types, status string) (Domain, error)
	Store(ctx context.Context, userDomain *Domain) (Domain, error)
	Update(ctx context.Context, userDomain *Domain) (Domain, error)
	Delete(ctx context.Context, userDomain *Domain) (Domain, error)
}
