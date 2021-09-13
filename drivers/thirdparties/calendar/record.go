package calendar

import (
	calendarUsecase "hungry-baby/businesses/calendar"

	"github.com/rs/xid"
	"google.golang.org/api/calendar/v3"
)

func FromDomain(domain *calendarUsecase.Domain) *calendar.Event {
	var calendarDomain *calendar.Event
	var attendees []*calendar.EventAttendee
	for _, a := range domain.Attendee {
		attendees = append(attendees, &calendar.EventAttendee{
			Email: a.Email,
		})
	}
	body := &calendar.Event{
		Summary:     domain.Title,
		Description: domain.Description,
		Start: &calendar.EventDateTime{
			DateTime: domain.StartAt,
		},
		End: &calendar.EventDateTime{
			DateTime: domain.EndAt,
		},
		Attendees: attendees,
	}

	if domain.CreateMeet {
		body.ConferenceData = &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId: xid.New().String(),
				ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
					Type: "hangoutsMeet",
				},
			},
		}
	}

	return calendarDomain
}

func ToDomain(rec *calendar.Event) calendarUsecase.Domain {
	var attendees []calendarUsecase.DomainAttendee
	for _, a := range rec.Attendees {
		attendees = append(attendees, calendarUsecase.DomainAttendee{
			Email: a.Email,
		})
	}

	return calendarUsecase.Domain{
		ID:          rec.Id,
		Title:       rec.Summary,
		Description: rec.Description,
		StartAt:     rec.Start.DateTime,
		EndAt:       rec.End.DateTime,
		Attendee:    attendees,
		MeetURL:     rec.HangoutLink,
	}
}
