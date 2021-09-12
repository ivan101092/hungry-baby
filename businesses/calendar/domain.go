package calendar

import (
	"context"

	"google.golang.org/api/calendar/v3"
)

// Domain ...
type Domain struct {
	Title       string
	Description string
	StartAt     string
	EndAt       string
	Attendee    []DomainAttendee
	CreateMeet  bool
}

// DomainAttendee ...
type DomainAttendee struct {
	Email string `json:"email" validate:"required"`
}

type Repository interface {
	FindAll(ctx context.Context, tokenString, search, startAt, endAt, pageToken string, limit int) (*calendar.Events, error)
	FindByID(ctx context.Context, tokenString, id string) (*calendar.Event, error)
	Add(ctx context.Context, tokenString string, body *calendar.Event) (*calendar.Event, error)
	Delete(ctx context.Context, tokenString, id string) error
}

type Usecase interface {
	FindAll(ctx context.Context, search, startAt, endAt, pageToken string, limit int) (*calendar.Events, error)
	FindByID(ctx context.Context, id string) (*calendar.Event, error)
	Store(ctx context.Context, calendarDomain *Domain) (*calendar.Event, error)
	Delete(ctx context.Context, id string) error
}
