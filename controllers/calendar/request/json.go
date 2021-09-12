package request

import "hungry-baby/businesses/calendar"

// Calendar ...
type Calendar struct {
	Title       string             `json:"title" validate:"required"`
	Description string             `json:"description" validate:"required"`
	StartAt     string             `json:"start_at" validate:"required"`
	EndAt       string             `json:"end_at" validate:"required"`
	Attendee    []CalendarAttendee `json:"attendee"`
	CreateMeet  bool               `json:"create_meet"`
}

// CalendarAttendee ...
type CalendarAttendee struct {
	Email string `json:"email" validate:"required"`
}

func (req *Calendar) ToDomain() *calendar.Domain {
	var attendee []calendar.DomainAttendee
	for _, a := range req.Attendee {
		attendee = append(attendee, calendar.DomainAttendee{
			Email: a.Email,
		})
	}

	return &calendar.Domain{
		Title:       req.Title,
		Description: req.Description,
		StartAt:     req.StartAt,
		EndAt:       req.EndAt,
		Attendee:    attendee,
		CreateMeet:  req.CreateMeet,
	}
}
