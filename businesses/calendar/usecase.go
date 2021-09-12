package calendar

import (
	"context"
	userCredentialBusiness "hungry-baby/businesses/userCredential"
	"hungry-baby/helpers/interfacepkg"
	"time"

	"github.com/rs/xid"
	"google.golang.org/api/calendar/v3"
)

type calendarUsecase struct {
	calendarRepository    Repository
	userCredentialUsecase userCredentialBusiness.Usecase
	contextTimeout        time.Duration
}

func NewCalendarUsecase(timeout time.Duration, repo Repository, userCredentialUsecase userCredentialBusiness.Usecase) Usecase {
	return &calendarUsecase{
		calendarRepository:    repo,
		userCredentialUsecase: userCredentialUsecase,
		contextTimeout:        timeout,
	}
}

func (uc *calendarUsecase) FindAll(ctx context.Context, search, startAt, endAt, pageToken string, limit int) (*calendar.Events, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userCredential, err := uc.userCredentialUsecase.FindByUserType(ctx, ctx.Value("userID").(int), "gmail", "true")
	if err != nil {
		return nil, err
	}

	res, err := uc.calendarRepository.FindAll(ctx, interfacepkg.Marshal(userCredential.RegistrationDetails), search, startAt, endAt, pageToken, limit)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *calendarUsecase) FindByID(ctx context.Context, id string) (*calendar.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userCredential, err := uc.userCredentialUsecase.FindByUserType(ctx, ctx.Value("userID").(int), "gmail", "true")
	if err != nil {
		return nil, err
	}

	res, err := uc.calendarRepository.FindByID(ctx, interfacepkg.Marshal(userCredential.RegistrationDetails), id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *calendarUsecase) Store(ctx context.Context, calendarDomain *Domain) (*calendar.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userCredential, err := uc.userCredentialUsecase.FindByUserType(ctx, ctx.Value("userID").(int), "gmail", "true")
	if err != nil {
		return nil, err
	}

	var attendees []*calendar.EventAttendee
	for _, a := range calendarDomain.Attendee {
		attendees = append(attendees, &calendar.EventAttendee{
			Email: a.Email,
		})
	}
	body := &calendar.Event{
		Summary:     calendarDomain.Title,
		Description: calendarDomain.Description,
		Start: &calendar.EventDateTime{
			DateTime: calendarDomain.StartAt,
		},
		End: &calendar.EventDateTime{
			DateTime: calendarDomain.EndAt,
		},
		Attendees: attendees,
	}

	if calendarDomain.CreateMeet {
		body.ConferenceData = &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId: xid.New().String(),
				ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
					Type: "hangoutsMeet",
				},
			},
		}
	}

	res, err := uc.calendarRepository.Add(ctx, interfacepkg.Marshal(userCredential.RegistrationDetails), body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *calendarUsecase) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userCredential, err := uc.userCredentialUsecase.FindByUserType(ctx, ctx.Value("userID").(int), "gmail", "true")
	if err != nil {
		return err
	}

	err = uc.calendarRepository.Delete(ctx, interfacepkg.Marshal(userCredential.RegistrationDetails), id)
	if err != nil {
		return err
	}

	return nil
}
