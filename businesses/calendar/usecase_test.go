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

// func TestStore(t *testing.T) {
// 	t.Run("test case 1, valid test", func(t *testing.T) {
// 		domain := calendar.Domain{
// 			ID:         1,
// 			Type:       "calendar",
// 			URL:        "calendar.jpg",
// 			FullURL:    "https://s3.hungrybaby.com/calendar.jpg",
// 			UserUpload: "1",
// 		}
// 		ctx := context.WithValue(context.Background(), "userID", 1)
// 		calendarRepository.On("Store", mock.Anything, &domain).Return(domain, nil).Once()
// 		minioRepository.On("GetCalendar", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/calendar.jpg", nil).Once()
// 		result, err := calendarUsecase.Store(ctx, &domain)

// 		assert.Equal(t, result, domain)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("test case 2, repository error", func(t *testing.T) {
// 		domain := calendar.Domain{
// 			ID:         1,
// 			Type:       "calendar",
// 			URL:        "calendar.jpg",
// 			FullURL:    "https://s3.hungrybaby.com/calendar.jpg",
// 			UserUpload: "1",
// 		}
// 		ctx := context.WithValue(context.Background(), "userID", 1)
// 		errNotFound := errors.New("(Repo) ID Not Found")
// 		calendarRepository.On("Store", mock.Anything, &domain).Return(calendar.Domain{}, errNotFound).Once()
// 		minioRepository.On("GetCalendar", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/calendar.jpg", nil).Once()
// 		result, err := calendarUsecase.Store(ctx, &domain)

// 		assert.Equal(t, result, calendar.Domain{})
// 		assert.Equal(t, err, errNotFound)
// 	})
// }

// // func TestStoreMinioErr(t *testing.T) {
// // 	t.Run("test case 3, minio error", func(t *testing.T) {
// // 		domain := calendar.Domain{
// // 			ID:         1,
// // 			Type:       "calendar",
// // 			URL:        "calendar.jpg",
// // 			FullURL:    "",
// // 			UserUpload: "1",
// // 		}
// // 		ctx := context.WithValue(context.Background(), "userID", 1)
// // 		errError := errors.New("(Minio) Error")
// // 		calendarRepository.On("Store", mock.Anything, mock.AnythingOfType("*calendar.Domain")).Return(domain, nil).Once()
// // 		minioRepository.On("GetCalendar", mock.AnythingOfType("string")).Return("", errError).Once()
// // 		result, err := calendarUsecase.Store(ctx, &domain)

// // 		assert.Equal(t, result, calendar.Domain{})
// // 		assert.Equal(t, errError, err)
// // 	})
// // }

// func TestUpload(t *testing.T) {
// 	t.Run("test case 3, repository error", func(t *testing.T) {
// 		errNotFound := errors.New("(Repo) ID Not Found")
// 		var calendarHeader *multipart.CalendarHeader
// 		domain := calendar.Domain{
// 			Type:       "calendar",
// 			URL:        "calendar.jpg",
// 			FullURL:    "https://s3.hungrybaby.com/calendar.jpg",
// 			UserUpload: "1",
// 		}
// 		ctx := context.WithValue(context.Background(), "userID", 1)
// 		minioRepository.On("Upload", mock.AnythingOfType("string"), mock.Anything).Return("calendar.jpg", nil).Once()
// 		domainStore := calendar.Domain{
// 			Type:       "calendar",
// 			URL:        "calendar.jpg",
// 			FullURL:    "",
// 			UserUpload: "1",
// 		}
// 		calendarRepository.On("Store", mock.Anything, &domainStore).Return(calendar.Domain{}, errNotFound).Once()
// 		minioRepository.On("GetCalendar", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/calendar.jpg", nil).Once()
// 		result, err := calendarUsecase.Upload(ctx, domain.Type, "", calendarHeader)

// 		assert.Equal(t, result, calendar.Domain{})
// 		assert.Equal(t, errNotFound, err)
// 	})

// 	t.Run("test case 1, valid test", func(t *testing.T) {
// 		var calendarHeader *multipart.CalendarHeader
// 		domain := calendar.Domain{
// 			Type:       "calendar",
// 			URL:        "calendar.jpg",
// 			FullURL:    "https://s3.hungrybaby.com/calendar.jpg",
// 			UserUpload: "1",
// 		}
// 		ctx := context.WithValue(context.Background(), "userID", 1)
// 		minioRepository.On("Upload", mock.AnythingOfType("string"), mock.Anything).Return("calendar.jpg", nil).Once()
// 		domainStore := calendar.Domain{
// 			Type:       "calendar",
// 			URL:        "calendar.jpg",
// 			FullURL:    "",
// 			UserUpload: "1",
// 		}
// 		calendarRepository.On("Store", mock.Anything, &domainStore).Return(domain, nil).Once()
// 		minioRepository.On("GetCalendar", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/calendar.jpg", nil).Once()
// 		result, err := calendarUsecase.Upload(ctx, domain.Type, "", calendarHeader)

// 		assert.Equal(t, result, domain)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("test case 2, minio error", func(t *testing.T) {
// 		errError := errors.New("(Minio) Error")
// 		var calendarHeader *multipart.CalendarHeader
// 		domain := calendar.Domain{
// 			Type:       "calendar",
// 			URL:        "calendar.jpg",
// 			FullURL:    "https://s3.hungrybaby.com/calendar.jpg",
// 			UserUpload: "1",
// 		}
// 		ctx := context.WithValue(context.Background(), "userID", 1)
// 		minioRepository.On("Upload", mock.AnythingOfType("string"), mock.Anything).Return("", errError).Once()
// 		domainStore := calendar.Domain{
// 			Type:       "calendar",
// 			URL:        "calendar.jpg",
// 			FullURL:    "",
// 			UserUpload: "1",
// 		}
// 		calendarRepository.On("Store", mock.Anything, &domainStore).Return(domainStore, nil).Once()
// 		minioRepository.On("GetCalendar", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/calendar.jpg", nil).Once()
// 		result, err := calendarUsecase.Upload(ctx, domain.Type, "", calendarHeader)

// 		assert.Equal(t, result, calendar.Domain{})
// 		assert.Equal(t, err, errError)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	t.Run("test case 1, valid test", func(t *testing.T) {
// 		domain := calendar.Domain{
// 			ID:         1,
// 			Type:       "calendar",
// 			URL:        "calendar.jpg",
// 			FullURL:    "https://s3.hungrybaby.com/calendar.jpg",
// 			UserUpload: "1",
// 		}
// 		ctx := context.WithValue(context.Background(), "userID", 1)
// 		calendarRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
// 		minioRepository.On("Delete", mock.AnythingOfType("string")).Return(nil).Once()
// 		calendarRepository.On("Delete", mock.Anything, mock.AnythingOfType("*calendar.Domain")).Return(domain, nil).Once()
// 		result, err := calendarUsecase.Delete(ctx, &domain)

// 		assert.Equal(t, result, domain)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("test case 2, repo error", func(t *testing.T) {
// 		errError := errors.New("(Repo) Error")
// 		domain := calendar.Domain{
// 			ID:         1,
// 			Type:       "calendar",
// 			URL:        "calendar.jpg",
// 			FullURL:    "https://s3.hungrybaby.com/calendar.jpg",
// 			UserUpload: "1",
// 		}
// 		ctx := context.WithValue(context.Background(), "userID", 1)
// 		calendarRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(calendar.Domain{}, errError).Once()
// 		minioRepository.On("Delete", mock.AnythingOfType("string")).Return(nil).Once()
// 		calendarRepository.On("Delete", mock.Anything, mock.AnythingOfType("*calendar.Domain")).Return(domain, nil).Once()
// 		result, err := calendarUsecase.Delete(ctx, &domain)

// 		assert.Equal(t, result, calendar.Domain{})
// 		assert.Equal(t, err, errError)
// 	})

// 	// t.Run("test case 3, repository error", func(t *testing.T) {
// 	// 	errNotFound := errors.New("(Repo) ID Not Found")
// 	// 	var calendarHeader *multipart.CalendarHeader
// 	// 	domain := calendar.Domain{
// 	// 		Type:       "calendar",
// 	// 		URL:        "calendar.jpg",
// 	// 		FullURL:    "https://s3.hungrybaby.com/calendar.jpg",
// 	// 		UserUpload: "1",
// 	// 	}
// 	// 	ctx := context.WithValue(context.Background(), "userID", 1)
// 	// 	minioRepository.On("Upload", mock.AnythingOfType("string"), mock.Anything).Return("calendar.jpg", nil).Once()
// 	// 	domainStore := calendar.Domain{
// 	// 		Type:       "calendar",
// 	// 		URL:        "calendar.jpg",
// 	// 		FullURL:    "",
// 	// 		UserUpload: "1",
// 	// 	}
// 	// 	calendarRepository.On("Store", mock.Anything, &domainStore).Return(calendar.Domain{}, errNotFound).Once()
// 	// 	minioRepository.On("GetCalendar", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/calendar.jpg", nil).Once()
// 	// 	result, err := calendarUsecase.Upload(ctx, domain.Type, "", calendarHeader)

// 	// 	assert.Equal(t, result, calendar.Domain{})
// 	// 	assert.Equal(t, errNotFound, err)
// 	// })
// }
