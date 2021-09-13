package calendar

import (
	"context"
)

// Domain ...
type Domain struct {
	ID          string
	Title       string
	Description string
	StartAt     string
	EndAt       string
	Attendee    []DomainAttendee
	CreateMeet  bool
	MeetURL     string
}

// DomainAttendee ...
type DomainAttendee struct {
	Email string `json:"email" validate:"required"`
}

type Repository interface {
	FindAll(ctx context.Context, tokenString, search, startAt, endAt, pageToken string, limit int) ([]Domain, error)
	FindByID(ctx context.Context, tokenString, id string) (Domain, error)
	Add(ctx context.Context, tokenString string, calendarDomain *Domain) (Domain, error)
	Delete(ctx context.Context, tokenString, id string) error
}

type Usecase interface {
	FindAll(ctx context.Context, search, startAt, endAt, pageToken string, limit int) ([]Domain, error)
	FindByID(ctx context.Context, id string) (Domain, error)
	Store(ctx context.Context, calendarDomain *Domain) (Domain, error)
	Delete(ctx context.Context, id string) error
}
