package calendar_test

import (
	"context"
	"errors"
	calendar "hungry-baby/businesses/calendar"
	userCredential "hungry-baby/businesses/userCredential"
	"hungry-baby/helpers/interfacepkg"
	calendarMock "hungry-baby/mocks/calendar"
	userCredentialMock "hungry-baby/mocks/userCredential"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			calendarRepository    calendarMock.Repository
			userCredentialUsecase userCredentialMock.Usecase
		)
		calendarUsecase := calendar.NewCalendarUsecase(2, &calendarRepository, &userCredentialUsecase)

		domain := []calendar.Domain{
			{
				ID:          "1",
				Title:       "calendar",
				Description: "desc",
				StartAt:     "2021-10-10",
				EndAt:       "2021-10-10",
				Attendee:    []calendar.DomainAttendee{},
				CreateMeet:  true,
				MeetURL:     "https://meet.com/google",
			},
		}
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, nil).Once()
		calendarRepository.On("FindAll", mock.Anything, interfacepkg.Marshal(userCredential.DomainRegistrationDetails{}),
			mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := calendarUsecase.FindAll(ctx, "", "", "", "", 1)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error credential", func(t *testing.T) {
		var (
			calendarRepository    calendarMock.Repository
			userCredentialUsecase userCredentialMock.Usecase
		)
		calendarUsecase := calendar.NewCalendarUsecase(2, &calendarRepository, &userCredentialUsecase)

		domain := []calendar.Domain{
			{
				ID:          "1",
				Title:       "calendar",
				Description: "desc",
				StartAt:     "2021-10-10",
				EndAt:       "2021-10-10",
				Attendee:    []calendar.DomainAttendee{},
				CreateMeet:  true,
				MeetURL:     "https://meet.com/google",
			},
		}
		errCredential := errors.New("Error Credential")
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, errCredential).Once()
		calendarRepository.On("FindAll", mock.Anything, mock.Anything,
			mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := calendarUsecase.FindAll(ctx, "", "", "", "", 1)

		assert.Equal(t, errCredential, err)
	})

	t.Run("test case 3, error calendar", func(t *testing.T) {
		var (
			calendarRepository    calendarMock.Repository
			userCredentialUsecase userCredentialMock.Usecase
		)
		calendarUsecase := calendar.NewCalendarUsecase(2, &calendarRepository, &userCredentialUsecase)

		errCredential := errors.New("Error Credential")
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, nil).Once()
		calendarRepository.On("FindAll", mock.Anything, mock.Anything,
			mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(nil, errCredential).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := calendarUsecase.FindAll(ctx, "", "", "", "", 1)

		assert.Equal(t, errCredential, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			calendarRepository    calendarMock.Repository
			userCredentialUsecase userCredentialMock.Usecase
		)
		calendarUsecase := calendar.NewCalendarUsecase(2, &calendarRepository, &userCredentialUsecase)

		domain := calendar.Domain{
			ID:          "1",
			Title:       "calendar",
			Description: "desc",
			StartAt:     "2021-10-10",
			EndAt:       "2021-10-10",
			Attendee:    []calendar.DomainAttendee{},
			CreateMeet:  true,
			MeetURL:     "https://meet.com/google",
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, nil).Once()
		calendarRepository.On("FindByID", mock.Anything, interfacepkg.Marshal(userCredential.DomainRegistrationDetails{}),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		result, err := calendarUsecase.FindByID(ctx, "1")

		assert.Equal(t, result, domain)
		assert.Nil(t, err)
	})

	t.Run("test case 2, error credential", func(t *testing.T) {
		var (
			calendarRepository    calendarMock.Repository
			userCredentialUsecase userCredentialMock.Usecase
		)
		calendarUsecase := calendar.NewCalendarUsecase(2, &calendarRepository, &userCredentialUsecase)

		domain := calendar.Domain{
			ID:          "1",
			Title:       "calendar",
			Description: "desc",
			StartAt:     "2021-10-10",
			EndAt:       "2021-10-10",
			Attendee:    []calendar.DomainAttendee{},
			CreateMeet:  true,
			MeetURL:     "https://meet.com/google",
		}
		errCredential := errors.New("Error Credential")
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, errCredential).Once()
		calendarRepository.On("FindByID", mock.Anything, mock.Anything,
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := calendarUsecase.FindByID(ctx, "1")

		assert.Equal(t, errCredential, err)
	})

	t.Run("test case 3, error calendar", func(t *testing.T) {
		var (
			calendarRepository    calendarMock.Repository
			userCredentialUsecase userCredentialMock.Usecase
		)
		calendarUsecase := calendar.NewCalendarUsecase(2, &calendarRepository, &userCredentialUsecase)

		errCredential := errors.New("Error Credential")
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, nil).Once()
		calendarRepository.On("FindByID", mock.Anything, mock.Anything,
			mock.AnythingOfType("string")).Return(calendar.Domain{}, errCredential).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := calendarUsecase.FindByID(ctx, "1")

		assert.Equal(t, errCredential, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			calendarRepository    calendarMock.Repository
			userCredentialUsecase userCredentialMock.Usecase
		)
		calendarUsecase := calendar.NewCalendarUsecase(2, &calendarRepository, &userCredentialUsecase)

		domain := calendar.Domain{
			ID:          "1",
			Title:       "calendar",
			Description: "desc",
			StartAt:     "2021-10-10",
			EndAt:       "2021-10-10",
			Attendee:    []calendar.DomainAttendee{},
			CreateMeet:  true,
			MeetURL:     "https://meet.com/google",
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, nil).Once()
		calendarRepository.On("Add", mock.Anything, interfacepkg.Marshal(userCredential.DomainRegistrationDetails{}),
			mock.AnythingOfType("*calendar.Domain")).Return(domain, nil).Once()
		result, err := calendarUsecase.Store(ctx, &domain)

		assert.Equal(t, result, domain)
		assert.Nil(t, err)
	})

	t.Run("test case 2, error credential", func(t *testing.T) {
		var (
			calendarRepository    calendarMock.Repository
			userCredentialUsecase userCredentialMock.Usecase
		)
		calendarUsecase := calendar.NewCalendarUsecase(2, &calendarRepository, &userCredentialUsecase)

		domain := calendar.Domain{
			ID:          "1",
			Title:       "calendar",
			Description: "desc",
			StartAt:     "2021-10-10",
			EndAt:       "2021-10-10",
			Attendee:    []calendar.DomainAttendee{},
			CreateMeet:  true,
			MeetURL:     "https://meet.com/google",
		}
		errCredential := errors.New("Error Credential")
		ctx := context.WithValue(context.Background(), "userID", 1)
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, errCredential).Once()
		calendarRepository.On("Add", mock.Anything, interfacepkg.Marshal(userCredential.DomainRegistrationDetails{}),
			mock.AnythingOfType("*calendar.Domain")).Return(domain, nil).Once()
		_, err := calendarUsecase.Store(ctx, &domain)

		assert.Equal(t, errCredential, err)
	})

	t.Run("test case 3, error repo", func(t *testing.T) {
		var (
			calendarRepository    calendarMock.Repository
			userCredentialUsecase userCredentialMock.Usecase
		)
		calendarUsecase := calendar.NewCalendarUsecase(2, &calendarRepository, &userCredentialUsecase)

		domain := calendar.Domain{
			ID:          "1",
			Title:       "calendar",
			Description: "desc",
			StartAt:     "2021-10-10",
			EndAt:       "2021-10-10",
			Attendee:    []calendar.DomainAttendee{},
			CreateMeet:  true,
			MeetURL:     "https://meet.com/google",
		}
		errRepo := errors.New("Error Repo")
		ctx := context.WithValue(context.Background(), "userID", 1)
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, nil).Once()
		calendarRepository.On("Add", mock.Anything, interfacepkg.Marshal(userCredential.DomainRegistrationDetails{}),
			mock.AnythingOfType("*calendar.Domain")).Return(calendar.Domain{}, errRepo).Once()
		_, err := calendarUsecase.Store(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			calendarRepository    calendarMock.Repository
			userCredentialUsecase userCredentialMock.Usecase
		)
		calendarUsecase := calendar.NewCalendarUsecase(2, &calendarRepository, &userCredentialUsecase)

		ctx := context.WithValue(context.Background(), "userID", 1)
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, nil).Once()
		calendarRepository.On("Delete", mock.Anything, interfacepkg.Marshal(userCredential.DomainRegistrationDetails{}),
			mock.AnythingOfType("string")).Return(nil).Once()
		err := calendarUsecase.Delete(ctx, "1")

		assert.Nil(t, err)
	})

	t.Run("test case 2, error credential", func(t *testing.T) {
		var (
			calendarRepository    calendarMock.Repository
			userCredentialUsecase userCredentialMock.Usecase
		)
		calendarUsecase := calendar.NewCalendarUsecase(2, &calendarRepository, &userCredentialUsecase)

		errCredential := errors.New("Error Credential")
		ctx := context.WithValue(context.Background(), "userID", 1)
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, errCredential).Once()
		calendarRepository.On("Delete", mock.Anything, interfacepkg.Marshal(userCredential.DomainRegistrationDetails{}),
			mock.AnythingOfType("string")).Return(nil).Once()
		err := calendarUsecase.Delete(ctx, "1")

		assert.Equal(t, errCredential, err)
	})

	t.Run("test case 3, error repo", func(t *testing.T) {
		var (
			calendarRepository    calendarMock.Repository
			userCredentialUsecase userCredentialMock.Usecase
		)
		calendarUsecase := calendar.NewCalendarUsecase(2, &calendarRepository, &userCredentialUsecase)

		errRepo := errors.New("Error Repo")
		ctx := context.WithValue(context.Background(), "userID", 1)
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, nil).Once()
		calendarRepository.On("Delete", mock.Anything, interfacepkg.Marshal(userCredential.DomainRegistrationDetails{}),
			mock.AnythingOfType("string")).Return(errRepo).Once()
		err := calendarUsecase.Delete(ctx, "1")

		assert.Equal(t, errRepo, err)
	})
}
