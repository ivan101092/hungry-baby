package calendar

import (
	"context"
	calendarBusiness "hungry-baby/businesses/calendar"
	"hungry-baby/drivers/thirdparties/google"
	"strings"

	"google.golang.org/api/calendar/v3"

	googlepkg "golang.org/x/oauth2/google"
)

// Credential ...
type Credential struct {
	Key         string
	RedirectURL string
}

func NewCalendar(key, redirectURL string) calendarBusiness.Repository {
	return &Credential{
		Key:         key,
		RedirectURL: redirectURL,
	}
}

// FindAll ...
func (cred *Credential) FindAll(ctx context.Context, tokenString, search, startAt, endAt, pageToken string, limit int) (res *calendar.Events, err error) {
	key := strings.Replace(cred.Key, "{redirect_url}", cred.RedirectURL, 1)
	b := []byte(key)
	config, err := googlepkg.ConfigFromJSON(b, google.CalendarScopes...)
	if err != nil {
		return res, err
	}
	client := google.GetClient(config, tokenString)

	srv, err := calendar.New(client)
	if err != nil {
		return res, err
	}

	eventListCall := srv.Events.List("primary").Q(search).ShowDeleted(false).PageToken(pageToken).MaxResults(int64(limit))
	if startAt != "" && endAt != "" {
		eventListCall = srv.Events.List("primary").Q(search).TimeMin(startAt).TimeMax(endAt).ShowDeleted(false).PageToken(pageToken).MaxResults(int64(limit))
	} else if startAt != "" {
		eventListCall = srv.Events.List("primary").Q(search).TimeMin(startAt).ShowDeleted(false).PageToken(pageToken).MaxResults(int64(limit))
	} else if endAt != "" {
		eventListCall = srv.Events.List("primary").Q(search).TimeMax(endAt).ShowDeleted(false).PageToken(pageToken).MaxResults(int64(limit))
	}
	res, err = eventListCall.Do()
	if err != nil {
		return res, err
	}

	return res, err
}

// FindByID ...
func (cred *Credential) FindByID(ctx context.Context, tokenString, id string) (res *calendar.Event, err error) {
	key := strings.Replace(cred.Key, "{redirect_url}", cred.RedirectURL, 1)
	b := []byte(key)
	config, err := googlepkg.ConfigFromJSON(b, google.CalendarScopes...)
	if err != nil {
		return res, err
	}
	client := google.GetClient(config, tokenString)

	srv, err := calendar.New(client)
	if err != nil {
		return res, err
	}

	res, err = srv.Events.Get("primary", id).Do()
	if err != nil {
		return res, err
	}

	return res, err
}

// Add ...
func (cred *Credential) Add(ctx context.Context, tokenString string, body *calendar.Event) (res *calendar.Event, err error) {
	key := strings.Replace(cred.Key, "{redirect_url}", cred.RedirectURL, 1)
	b := []byte(key)
	config, err := googlepkg.ConfigFromJSON(b, google.CalendarScopes...)
	if err != nil {
		return res, err
	}
	client := google.GetClient(config, tokenString)

	srv, err := calendar.New(client)
	if err != nil {
		return res, err
	}

	var converenceDataVersion int64
	if body.ConferenceData != nil {
		converenceDataVersion = 1
	}
	res, err = srv.Events.Insert("primary", body).ConferenceDataVersion(converenceDataVersion).Do()
	if err != nil {
		return res, err
	}

	return res, err
}

// Delete ...
func (cred *Credential) Delete(ctx context.Context, tokenString, id string) (err error) {
	key := strings.Replace(cred.Key, "{redirect_url}", cred.RedirectURL, 1)
	b := []byte(key)
	config, err := googlepkg.ConfigFromJSON(b, google.CalendarScopes...)
	if err != nil {
		return err
	}
	client := google.GetClient(config, tokenString)

	srv, err := calendar.New(client)
	if err != nil {
		return err
	}

	err = srv.Events.Delete("primary", id).Do()
	if err != nil {
		return err
	}

	return err
}
