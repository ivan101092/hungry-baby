package response

import (
	"hungry-baby/businesses/calendar"
)

// Calendar ...
type Calendar struct {
	ID          string             `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	StartAt     string             `json:"start_at"`
	EndAt       string             `json:"end_at"`
	Attendee    []CalendarAttendee `json:"attendee"`
	MeetURL     string             `json:"meet_url"`
}

// CalendarAttendee ...
type CalendarAttendee struct {
	Email string `json:"email"`
}

func FromDomain(domain calendar.Domain) Calendar {
	var attendee []CalendarAttendee
	for _, a := range domain.Attendee {
		attendee = append(attendee, CalendarAttendee{
			Email: a.Email,
		})
	}
	return Calendar{
		ID:          domain.ID,
		Title:       domain.Title,
		Description: domain.Description,
		StartAt:     domain.StartAt,
		EndAt:       domain.EndAt,
		Attendee:    attendee,
		MeetURL:     domain.MeetURL,
	}
}
